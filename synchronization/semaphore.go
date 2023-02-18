package synchronization

type Semaphore interface {
	Wait()
	Signal()
}

type semaphore struct {
	semChan chan struct{}
}

func NewSemaphore(val int) Semaphore {
	return &semaphore{
		make(chan struct{}, val),
	}
}

func (sem *semaphore) Wait() {
	sem.semChan <- struct{}{}
}

func (sem *semaphore) Signal() {
	<-sem.semChan
}
