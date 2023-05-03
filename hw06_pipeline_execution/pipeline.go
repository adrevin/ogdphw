package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	p := chanProxy(in, done)
	for _, stage := range stages {
		p = chanProxy(stage(p), done)
	}
	return p
}

func chanProxy(in In, done In) Out {
	out := make(Bi)
	go func() {
		for v := range in {
			select {
			case <-done:
				close(out)
				return
			default:
				out <- v
			}
		}
		close(out)
	}()
	return out
}
