package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

type handler struct {
	out Out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if in == nil {
		panic("in is nil")
	}
	if stages == nil {
		panic("stages is nil")
	}
	h := &handler{out: chanProxy(in, done)}
	for _, stage := range stages {
		h.out = chanProxy(stage(h.out), done)
	}
	return h.out
}

func chanProxy(in In, done In) Out {
	out := make(Bi)
	go func() {
		for {
			select {
			case <-done:
				close(out)
				return
			case v, ok := <-in:
				if ok {
					out <- v
				} else {
					close(out)
					return
				}
			}
		}
	}()
	return out
}
