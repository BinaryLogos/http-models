package httpmodels

// PromptRequest is the request body for a prompt request in the service promt-engine
type PromptRequest struct {
	// CategoryID is the category id used for a prompt generation
	CategoryID string `json:"category_id"`
	// PromptData is the custom data for a prompt generation
	PromptData PromptData `json:"prompt_data"`
}

// PromptData is the custom data for a prompt generation
type PromptData struct {
	Desciption      string   `json:"description"`
	InformationLine []string `json:"information_line"`
}

// PromptResponse is the response body for a prompt request in the service promt-engine
type PromptResponse struct {
	// Prompts that are generated
	Prompt []string `json:"prompts"`
}
