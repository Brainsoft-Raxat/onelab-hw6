package main

import (
	"errors"
	"fmt"
	"github.com/Brainsoft-Raxat/onelab-hw6/part2"
)

func NotOk() error {
	return errors.New("some error")
}

func Ok() error {
	return nil
}

func main() {
	testCase1 := make([]func() error, 10)
	E1 := 4

	for i, _ := range testCase1 {
		if i < E1 {
			testCase1[i] = NotOk
		} else {
			testCase1[i] = Ok
		}
	}

	err := part2.Execute(testCase1, E1)
	fmt.Println(err)
}
