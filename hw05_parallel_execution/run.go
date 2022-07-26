package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// increment and check if maxError exceeded.
func limitedIncrement(m *sync.Mutex, errorsTotal *int, maxError int) bool {
	m.Lock()
	defer m.Unlock()
	*errorsTotal++
	return *errorsTotal < maxError
}

func worker(w *sync.WaitGroup, m *sync.Mutex, in chan Task, errorsTotal *int, maxError int) {
	defer w.Done()
	for task := range in {
		if err := task(); err != nil {
			if !limitedIncrement(m, errorsTotal, maxError) {
				return
			}
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// todo - m < 0

	w := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	w.Add(n) // no more than n tasks
	inChannels := make(chan Task, len(tasks))

	// push tasks to channels
	for i := 0; i < len(tasks); i++ {
		inChannels <- tasks[i]
	}
	close(inChannels)

	// start workers
	var errorsTotal int
	for i := 0; i < n; i++ {
		go worker(w, mut, inChannels, &errorsTotal, m)
	}
	w.Wait()

	// return
	var result error
	if errorsTotal >= m {
		result = ErrErrorsLimitExceeded
	}
	return result
}
