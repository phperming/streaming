package taskrunner

type Runner struct {
	Controller controlChan
	Error controlChan
	Data dataChan
	dataSize int
	longLived bool
	Dispatcher fn
	Executor fn
}

func NewRunner(size int,longLived bool,d fn, e fn) *Runner  {
	return &Runner{
		Controller: make(chan string,1),
		Error: make(chan string,1),
		Data: make(chan interface{},size),
		longLived: longLived,
		Dispatcher: d,
		Executor: e,
	}
}

func (r *Runner)startDispatch()  {
	defer func() {
		if !r.longLived {
			close(r.Controller)
			close(r.Error)
			close(r.Data)
		}
	}()

	for  {
		select {
		case c := <- r.Controller :
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}

			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e := <- r.Error:
			if e == CLOSE {
				return
			}
		default:

		}
	}

}

func (r *Runner)StartAll()  {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}


