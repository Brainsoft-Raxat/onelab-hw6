package part2

import (
	"context"
	"errors"
	"sync"
)

var ExecuteError = errors.New("too much errors")

func Execute(tasks []func() error, E int) error {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	//var wg sync.WaitGroup
	var mu sync.Mutex
	//totalTasks := len(tasks)
	ch := make(chan error)
	for _, task := range tasks {
		f := task
		go func(ch chan error) {
			mu.Lock()
			defer mu.Unlock()
			err := f()
			ch <- err
		}(ch)

	}

	tasksCount := 0
	errCount := 0
	for {
		select {
		case <-ctx.Done():
			return ExecuteError
		case err, ok := <-ch:
			if !ok {
				break
			}
			if err != nil {
				errCount++
			}
			tasksCount++
			if errCount >= E {
				cancelFunc()
				continue
			}
			if tasksCount >= len(tasks) {
				return nil
			}
		}
	}
}
