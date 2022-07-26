package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// increment and check if maxError exceeded.
func limitedIncrement(m *sync.Mutex, errorsTotal *int, maxErrorLimit int) bool {
	m.Lock()
	defer m.Unlock()
	*errorsTotal++
	if maxErrorLimit > 0 {
		return *errorsTotal < maxErrorLimit
	}
	return true
}

func worker(w *sync.WaitGroup, m *sync.Mutex, in chan Task, errorsTotal *int, maxErrorLimit int) {
	defer w.Done()
	for task := range in {
		if err := task(); err != nil {
			if !limitedIncrement(m, errorsTotal, maxErrorLimit) {
				return
			}
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// handle m <= 0 input value
	maxErrorLimit := m
	if m <= 0 {
		maxErrorLimit = -1
	}

	// init sync primitives
	mut := &sync.Mutex{}
	w := &sync.WaitGroup{}
	w.Add(n) // no more than n tasks

	// push tasks to channels
	inChannels := make(chan Task, len(tasks))
	for i := 0; i < len(tasks); i++ {
		inChannels <- tasks[i]
	}
	close(inChannels)

	// start workers
	var errorsTotal int
	for i := 0; i < n; i++ {
		go worker(w, mut, inChannels, &errorsTotal, maxErrorLimit)
	}
	w.Wait()

	// return
	var result error
	if maxErrorLimit > 0 && errorsTotal >= maxErrorLimit {
		result = ErrErrorsLimitExceeded
	}
	return result
}
