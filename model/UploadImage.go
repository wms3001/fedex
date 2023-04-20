package model

type UploadImage struct {
	Document struct {
		ReferenceId string `json:"referenceId"`
		Name        string `json:"name"`
		ContentType string `json:"contentType"`
		Meta        struct {
			ImageType  string `json:"imageType"`
			ImageIndex string `json:"imageIndex"`
		} `json:"meta"`
	} `json:"document"`
	Rules struct {
		WorkflowName string `json:"workflowName"`
	} `json:"rules"`
}
