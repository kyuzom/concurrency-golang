package main

import (
	"fmt"
	"time"
)

func generate(data string) <-chan string {
	chn := make(chan string)
	go func() {
		for {
			chn <- data
			time.Sleep(time.Duration(100 * time.Millisecond))
		}
	}()
	return chn
}

type Processor struct {
	jobChannel chan string
	done       chan *Worker
	workers    []*Worker
}

type Worker struct {
	name string
}

func getProcessor() *Processor {
	p := &Processor{
		jobChannel: make(chan string),
		done:       make(chan *Worker),
		workers:    make([]*Worker, 5),
	}
	for i := 0; i < 5; i++ {
		w := &Worker{
			name: fmt.Sprintf("<Worker - %d>", i),
		}
		p.workers[i] = w
	}
	p.startProcess()
	return p
}

func (p *Processor) startProcess() {
	go func() {
		for {
			select {
			default:
				if len(p.workers) > 0 {
					w := p.workers[0]
					p.workers = p.workers[1:]
					w.processJob(<-p.jobChannel, p.done)
				}
			case w := <-p.done:
				p.workers = append(p.workers, w)
			}
		}
	}()
}

func (w *Worker) processJob(data string, done chan *Worker) {
	go func() {
		fmt.Println("Working on data ", data, w.name)
		time.Sleep(time.Duration(2000 * time.Millisecond))
		done <- w
	}()
}

func (p *Processor) postJob(jobs <-chan string) {
	p.jobChannel <- <-jobs
}

func main() {
	source := generate("data string")
	process := getProcessor()
	for i := 0; i < 12; i++ {
		process.postJob(source)
	}
}
