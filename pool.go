package pool

import "sync"

type Pool struct {
	group *sync.WaitGroup
	mut   *sync.RWMutex
	queue chan bool
	size  int
}

func New(concurrency int) Pool {
	p := Pool{
		group: &sync.WaitGroup{},
		mut:   &sync.RWMutex{},
		queue: make(chan bool, concurrency),
	}
	return p
}

func (p Pool) Add() {
	p.queue <- true
	p.mut.Lock()
	p.size += 1
	defer p.mut.Unlock()
	p.group.Add(1)
}

func (p Pool) Done() {
	<-p.queue
	p.group.Done()
}

func (p Pool) Size() int {
	p.mut.RLock()
	defer p.mut.RUnlock()
	return p.size
}

func (p Pool) QueueSize() int {
	return len(p.queue)
}

func (p Pool) Wait() {
	p.group.Wait()
}
