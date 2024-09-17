package main

import (
	"log"

	"github.com/cu8code/9tail/api"
	"github.com/cu8code/9tail/internal/block"
	"github.com/cu8code/9tail/internal/scheduler"
	"github.com/cu8code/9tail/internal/workflow"
	"github.com/nats-io/nats.go"
)

func main() {
	log.Println("Connecting to NATS server...")
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to NATS server")

	defer func() {
		log.Println("Closing NATS connection...")
		nc.Close()
		log.Println("NATS connection closed")
	}()

	log.Println("Initializing workflow manager...")
	wm := workflow.NewManager(nc)
	log.Println("Workflow manager initialized")

	log.Println("Initializing scheduler...")
	s := scheduler.NewScheduler(nc)
	log.Println("Scheduler initialized")

	log.Println("Initializing block registry...")
	registry := block.NewBlockRegistry(nc, "block.registry")
	log.Println("Block registry initialized")

	log.Println("Registering blocks...")
	registry.RegisterBlock("fetch", "block.fetch.execute")
	registry.RegisterBlock("html-scraper", "block.html-scraper.execute")
	registry.RegisterBlock("hook", "block.hook.execute")
	registry.RegisterBlock("response", "block.response.execute")
	log.Println("Blocks registered")

	log.Println("Initializing executor...")
	executor := block.NewExecutor(nc, registry)
	log.Println("Executor initialized")

	log.Println("Initializing workflow context publisher...")
	workflowContext, err := workflow.NewContextPublisher(nc)
	if err != nil {
		log.Println("Error initializing workflow context publisher:", err)
		panic(err)
	}
	log.Println("Workflow context publisher initialized")

	// Start the components
	log.Println("Starting block registry...")
	registry.Start()
	log.Println("Starting workflow manager...")
	go wm.Start()
	log.Println("Starting scheduler...")
	go s.Start()
	log.Println("Starting workflow context publisher...")
	go workflowContext.Start()
	log.Println("Starting executor...")
	go executor.Start()
	log.Println("Starting HTTP server...")
	go api.StartHTTPServer(nc)

	log.Println("HTTP Server listening on port 8080")

	// Keep the main goroutine alive
	log.Println("Main goroutine waiting...")
	select {}
}
