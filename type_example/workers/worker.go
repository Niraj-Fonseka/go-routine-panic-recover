package workers

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

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

func (w *Worker) GetError() error {
	return w.Err
}

func (w *Worker) GetWorkerID() string {
	return w.ID
}

func (w *Worker) GetSleepDuration() time.Duration {
	return w.Duration
}

func (w *Worker) Work(worker chan<- *Worker) (err error) {
	fmt.Printf("\033[37mStarting Worker :\033[0m \033[34m%s\033[0m \n\n", w.GetWorkerID())

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				w.Err = err
			} else {
				w.Err = fmt.Errorf("error : %v", r)
			}
		} else {
			w.Err = err
		}
		worker <- w
	}()

	for {
		rand.Seed(time.Now().UnixNano())
		b := big.NewInt(int64(rand.Intn(100)))

		fmt.Printf("\033[34m%s\033[0m \033[37mdoing work ..\033[0m\n\n", w.GetWorkerID())
		if b.ProbablyPrime(0) {
			panic(fmt.Sprintf("random %d is prime \n", b))
		}

		time.Sleep(w.GetSleepDuration())
	}
}
