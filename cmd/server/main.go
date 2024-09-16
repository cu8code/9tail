package main

import (
	"log"

	"github.com/cu8code/9tail/internal/block"
	"github.com/cu8code/9tail/internal/scheduler"
	"github.com/cu8code/9tail/internal/workflow"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	wm := workflow.NewManager(nc)
	s := scheduler.NewScheduler(nc)
	be := block.NewExecutor(nc)

	// Start the components
	go wm.Start()
	go s.Start()
	go be.Start()

	log.Println("NATS-based Workflow System started")

	// Keep the main goroutine alive
	select {}
}
