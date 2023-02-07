package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	worker_one := NewWorker("one", 2*time.Second)
	worker_two := NewWorker("two", 5*time.Second)

	workers := make(chan WorkerInterface, 2)

	go worker_one.Work(workers)
	go worker_two.Work(workers)

	for w := range workers {
		fmt.Printf("panic happened in the worker : %s\n", w.GetWorkerID())

		go w.(*Worker).Work(workers)
	}

}

type WorkerInterface interface {
	GetWorkerID() string
	Work(worker chan<- WorkerInterface) (err error)
}

type Worker struct {
	ID       string
	Err      error
	Duration time.Duration
}

func NewWorker(id string, sleep time.Duration) *Worker {
	return &Worker{
		ID:       id,
		Duration: sleep,
	}
}

func (w *Worker) GetWorkerID() string {
	return w.ID
}

func (w *Worker) GetSleepDuration() time.Duration {
	return w.Duration
}

func (w *Worker) Work(worker chan<- WorkerInterface) (err error) {
	fmt.Printf("Starting Worker : %s \n", w.GetWorkerID())
	rand.Seed(time.Now().Unix())

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				w.Err = err
			} else {
				w.Err = fmt.Errorf("Panic happened with %v", r)
			}
		} else {
			w.Err = err
		}
		worker <- w
	}()

	for {
		r := rand.Intn(100)

		fmt.Printf("worker %s doing work ..\n", w.GetWorkerID())
		if r%2 == 0 {
			panic(fmt.Errorf("error happened because the random number was : %d \n", r))
		}

		time.Sleep(w.GetSleepDuration())
	}
}
