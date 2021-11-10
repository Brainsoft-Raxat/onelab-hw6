package part1

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NotOk() error {
	return errors.New("some error")
}

func Ok() error {
	return nil
}

func TestExecute(t *testing.T) {
	testCase1 := make([]func() error, 10)
	E1 := 4
	testCase2 := make([]func() error, 10)
	E2 := 4

	for i, _ := range testCase1 {
		if i < E1 {
			testCase1[i] = NotOk
		} else {
			testCase1[i] = Ok
		}
	}

	for i, _ := range testCase2 {
		testCase2[i] = Ok
	}

	t.Run("test function Execute if it returns err", func(t *testing.T) {
		assert.Equal(t, Execute(testCase1, E1), ExecuteError)
	})

	t.Run("test function Execute if it returns nil", func(t *testing.T) {
		assert.Equal(t, Execute(testCase2, E2), nil)
	})

}
