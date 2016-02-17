package pool

import "sync"

type Pool struct {
	group *sync.WaitGroup
	queue chan bool
	size  int
}

func New(concurrency int) Pool {
	p := Pool{
		group: &sync.WaitGroup{},
		queue: make(chan bool, concurrency),
	}
	return p
}

func (p Pool) Add() {
	p.queue <- true
	p.size += 1
	p.group.Add(1)
}

func (p Pool) Done() {
	<-p.queue
	p.group.Done()
}

func (p Pool) Size() int {
	return p.size
}

func (p Pool) QueueSize() int {
	return len(p.queue)
}

func (p Pool) Wait() {
	p.group.Wait()
}
