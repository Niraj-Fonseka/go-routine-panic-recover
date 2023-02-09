package workers

type WorkerInterface interface {
	GetWorkerID() string
	Work(worker chan<- WorkerInterface) (err error)
	GetError() error
}
