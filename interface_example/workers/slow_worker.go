package workers

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

type SlowWorker struct {
	ID       string
	Err      error
	Duration time.Duration
}

func NewSlowWorker(id string, sleep time.Duration) *SlowWorker {
	return &SlowWorker{
		ID:       id,
		Duration: sleep,
	}
}

func (w *SlowWorker) GetWorkerID() string {
	return w.ID
}

func (w *SlowWorker) GetSleepDuration() time.Duration {
	return w.Duration
}

func (w *SlowWorker) Work(worker chan<- WorkerInterface) (err error) {
	fmt.Printf("\033[37mStarting Worker :\033[0m \033[34m%s\033[0m \n\n", w.GetWorkerID())

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				w.Err = err
			} else {
				w.Err = fmt.Errorf("PANIC happened : %v \n", r)
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
			panic(fmt.Sprintf("error happened because the random %d is prime \n", b))
		}

		time.Sleep(10 * time.Second)
	}
}
