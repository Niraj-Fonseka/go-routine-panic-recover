package main

import (
	"fmt"
	"go-routine-panic-recover/interface_example/workers"
	"time"
)

func main() {

	hard_worker_one := workers.NewHardWorker("hard_worker_1", 2*time.Second)
	slow_worker_one := workers.NewSlowWorker("slow_worker_1", 5*time.Second)

	wrks := make(chan workers.WorkerInterface, 2)

	go hard_worker_one.Work(wrks)
	go slow_worker_one.Work(wrks)

	for w := range wrks {
		fmt.Printf("\033[31m---------------- PANIC happened in worker : \033[0m\033[34m%s\033[0m\033[31m because %s\033[0m\n", w.GetWorkerID(), w.GetError().Error())

		go w.Work(wrks)

	}

}
