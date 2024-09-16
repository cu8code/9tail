package models

type BlockControlObject struct {
    Type         string        `json:"type"`        // Block type ('fetch', 'html-scraper', etc.)
    Input        interface{}   `json:"input"`       // Input for the block
    Output       interface{}   `json:"output"`      // Output from the block
    Status       string        `json:"status"`      // Block status ('PENDING', 'RUNNING', etc.)
    Retries      int           `json:"retries"`     // Retry count in case of failure
    Dependencies []string      `json:"dependencies"`// Block dependencies
}

