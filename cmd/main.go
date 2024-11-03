package main

import (
	"github.com/codewithkushal/go-concurrency-lab/src/executor"
	bakeryorderprocessing "github.com/codewithkushal/go-concurrency-lab/src/level1/bakery_order_processing"
)

func main() {
	executors := []executor.Executor{
		// level 1 - problem 1
		bakeryorderprocessing.NewBakeryOrderProcessing(),
	}
	for _, executor := range executors {
		executor.Execute()
	}
}
