package main

import (
	"go-routine-panic-recover/panic_example/workers"
	"time"
)

func main() {

	worker_one := workers.NewWorker("worker_1", 2*time.Second)

	go worker_one.Work()

	select {} // block forever
}
