package pool

import "sync"

type Pool struct {
	group *sync.WaitGroup
	mut   *sync.RWMutex
	queue chan struct{}
	errs  chan error
	err   error
	size  int
}

func New(concurrency int) *Pool {
	p := &Pool{
		group: &sync.WaitGroup{},
		mut:   &sync.RWMutex{},
		queue: make(chan struct{}, concurrency),
		errs:  make(chan error),
	}
	go p.listen()
	return p
}

func (p *Pool) listen() {
	for {
		err := <-p.errs
		if err != nil {
			p.err = err
		}
	}
}

func (p *Pool) Add() {
	p.queue <- struct{}{}
	p.mut.Lock()
	p.size += 1
	defer p.mut.Unlock()
	p.group.Add(1)
}

func (p *Pool) Done(err error) {
	<-p.queue
	p.errs <- err
	p.group.Done()
}

func (p *Pool) Size() int {
	p.mut.RLock()
	defer p.mut.RUnlock()
	return p.size
}

func (p *Pool) QueueSize() int {
	return len(p.queue)
}

func (p *Pool) Error() error {
	return p.err
}

func (p *Pool) Wait() {
	p.group.Wait()
}
