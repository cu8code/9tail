package execution

import (
	"fmt"

	"github.com/cu8code/9tail/pkg/models"
)

type WorkflowManager struct{}

func (wm *WorkflowManager) StartWorkflow(wcb *models.WorkflowControlBlock) {
	fmt.Println("Starting workflow: ", wcb.ID)
	for i := wcb.CurrentBlockIndex; i < len(wcb.Blocks); i++ {
		block := wcb.Blocks[i]
		if err := wm.ExecuteBlock(&block); err != nil {
			fmt.Println("Block failed:", block.Type, err)
			// Handle block failure (retry, skip, etc.)
		}
		wcb.CurrentBlockIndex = i
	}
	wcb.State = "COMPLETED"
	fmt.Println("Workflow completed:", wcb.ID)
}

func (wm *WorkflowManager) ExecuteBlock(bco *models.BlockControlObject) error {
	fmt.Println("Executing block:", bco.Type)
	// Process block logic based on its type (fetch, scraper, etc.)
	switch bco.Type {
	case "fetch":
		// Call fetch logic
	case "html-scraper":
		// Call scraper logic
	default:
		return fmt.Errorf("unknown block type: %s", bco.Type)
	}
	bco.Status = "COMPLETED"
	return nil
}
