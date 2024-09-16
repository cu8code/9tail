package main

import (
	"fmt"

	"github.com/cu8code/9tail/internal/execution"
	"github.com/cu8code/9tail/pkg/models"
)

func main() {
	fmt.Println("Starting Workflow System")

	// Sample workflow setup
	wcb := &models.WorkflowControlBlock{
		ID:    "workflow_1",
		Type:  "html-scraper",
		State: "READY",
		Blocks: []models.BlockControlObject{
			{Type: "fetch", Status: "PENDING"},
			{Type: "html-scraper", Status: "PENDING"},
		},
		Priority: 1,
	}

	scheduler := execution.Scheduler{}
	scheduler.Schedule([]*models.WorkflowControlBlock{wcb})
}
