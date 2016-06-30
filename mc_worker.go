package main

import (
	"fmt"
)

func mcWorker(s *MCService, stream *chan *MCJob) {
	fmt.Println("Worker...")
	for {
		j := <-*stream
		j.Response = mcGet(s, j.Path)
		go j.Done(j)
	}
}
