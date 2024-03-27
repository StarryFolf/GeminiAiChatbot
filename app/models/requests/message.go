package requests

type CreateMessageRequest struct {
	Data struct {
		Content        string `json:"content" validate:"required"`
		Sender         string `json:"sender" validate:"required"`
		ConversationId string `json:"conversationId"`
	} `json:"data"`
}
