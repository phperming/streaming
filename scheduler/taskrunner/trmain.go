package taskrunner

import "time"

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration,r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval),
		runner: r,
	}
}

func (w *Worker)startWork()  {
	for  {
		select {
		case <- w.ticker.C :
			go w.runner.StartAll()
		}
	}
}

func Start()  {
	r := NewRunner(3,true,VideoClearDispatcher,VideoClearExecutor)
	w := NewWorker(3,r)
	go w.startWork()
}