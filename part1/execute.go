package part1

import (
	"errors"
	"sync"
)

var ExecuteError = errors.New("too much errors")

func Execute(tasks []func() error, E int) error {
	ch := make(chan error, len(tasks))
	var wg sync.WaitGroup
	for _, task := range tasks {
		f := task
		wg.Add(1)
		go func(ch chan error) {
			defer wg.Done()
			err := f()
			if err != nil {
				ch <- err
			}
		}(ch)
	}
	wg.Wait()
	close(ch)

	counter := 0
	for err := range ch {
		if err != nil {
			counter++
		}
	}

	if counter >= E {
		return ExecuteError
	}

	return nil
}
