package main

import (
	"fmt"
	"go-routine-panic-recover/type_example/workers"
	"time"
)

func main() {

	worker_one := workers.NewWorker("worker_1", 2*time.Second)
	worker_two := workers.NewWorker("worker_1", 5*time.Second)

	wrks := make(chan *workers.Worker, 2)

	go worker_one.Work(wrks)
	go worker_two.Work(wrks)

	for w := range wrks {
		fmt.Printf("\033[31m---------------- PANIC happened in worker : \033[0m\033[34m%s\033[0m\033[31m because %s\033[0m\n", w.GetWorkerID(), w.GetError().Error())

		fmt.Printf("\033[32m-------------\033[0m \033[34m%s\033[0m \033[32mrecovering ...\033[0m \n", w.GetWorkerID())
		go w.Work(wrks)
	}

}
