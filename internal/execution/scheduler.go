package execution

import (
	"fmt"

	"github.com/cu8code/9tail/pkg/models"
)

type Scheduler struct{}

func (s *Scheduler) Schedule(workflows []*models.WorkflowControlBlock) {
	fmt.Println("Starting scheduling of workflows")
	for _, workflow := range workflows {
		if workflow.State == "READY" || workflow.State == "BLOCKED" {
			s.ExecuteWorkflow(workflow)
		}
	}
}

func (s *Scheduler) ExecuteWorkflow(wcb *models.WorkflowControlBlock) {
	fmt.Println("Executing workflow:", wcb.ID)
	wm := WorkflowManager{}
	wm.StartWorkflow(wcb)
}
