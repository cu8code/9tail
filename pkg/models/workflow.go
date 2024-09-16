package models

import "time"

type WorkflowControlBlock struct {
    ID                string                `json:"id"`                  // Unique workflow ID
    Type              string                `json:"type"`                // Workflow type ('fetch', 'html-scraper', etc.)
    State             string                `json:"state"`               // Workflow state ('READY', 'RUNNING', etc.)
    Blocks            []BlockControlObject  `json:"blocks"`              // List of blocks in the workflow
    Input             interface{}           `json:"input"`               // Input to the first block
    Output            interface{}           `json:"output"`              // Final output after all blocks execution
    CurrentBlockIndex int                   `json:"current_block_index"` // Index of the current block
    Dependencies      []string              `json:"dependencies"`        // List of workflow dependencies
    Priority          int                   `json:"priority"`            // Workflow priority level
    CreatedAt         time.Time             `json:"created_at"`          // Time of workflow creation
    UpdatedAt         time.Time             `json:"updated_at"`          // Last updated time
}

