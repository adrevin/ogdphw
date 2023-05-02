package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type handler struct {
	taskIndex       int32
	maxTaskIndex    int32
	errorsNumber    int32
	wg              sync.WaitGroup
	maxErrorsNumber int32
}

func Run(tasks []Task, n, m int) error {
	h := &handler{
		taskIndex:       -1,
		maxTaskIndex:    int32(len(tasks) - 1),
		errorsNumber:    0,
		maxErrorsNumber: int32(m),
	}

	for i := 0; i < n; i++ {
		h.wg.Add(1)
		go work(&tasks, h)
	}

	h.wg.Wait()
	// Значение m <= 0 трактуется как знак игнорировать ошибки в принципе
	if m <= 0 {
		return nil
	}

	if h.errorsNumber >= h.maxErrorsNumber {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func work(tasks *[]Task, h *handler) {
	defer h.wg.Done()
	for {
		index := atomic.AddInt32(&h.taskIndex, 1)
		if index > h.maxTaskIndex {
			return
		}
		if h.maxErrorsNumber > 0 && h.errorsNumber >= h.maxErrorsNumber {
			return
		}
		if (*tasks)[index]() != nil {
			atomic.AddInt32(&h.errorsNumber, 1)
		}
	}
}
