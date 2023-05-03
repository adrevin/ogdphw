package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type handler struct {
	errorsNumber    int32
	wg              sync.WaitGroup
	maxErrorsNumber int32
	tasks           chan Task
}

func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		panic("n must be greater than 0")
	}

	h := &handler{
		maxErrorsNumber: int32(m),
		tasks:           make(chan Task, len(tasks)),
	}

	for i := 0; i < len(tasks); i++ {
		h.tasks <- tasks[i]
	}
	close(h.tasks)

	for i := 0; i < n; i++ {
		h.wg.Add(1)
		go work(h)
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

func work(h *handler) {
	defer h.wg.Done()
	for task := range h.tasks {
		if h.maxErrorsNumber > 0 && atomic.LoadInt32(&h.errorsNumber) >= h.maxErrorsNumber {
			return
		}
		if task() != nil {
			atomic.AddInt32(&h.errorsNumber, 1)
		}
	}
}
