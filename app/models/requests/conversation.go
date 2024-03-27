package requests

type CreateConversationRequest struct {
	Data struct {
		Name           string                  `json:"name" validate:"required"`
		SampleMessages []*CreateMessageRequest `json:"sampleMessages" validate:"required"`
	} `json:"data"`
}
