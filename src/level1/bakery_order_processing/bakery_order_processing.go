package bakeryorderprocessing

import "github.com/codewithkushal/go-concurrency-lab/src/executor"

type BakeryOrderProcessing struct {
}

func NewBakeryOrderProcessing() executor.Executor {
	return &BakeryOrderProcessing{}
}

func (bop *BakeryOrderProcessing) Execute() {

}
