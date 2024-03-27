package dto

import (
	"fiber/app/models/entities"
	"github.com/google/uuid"
	"time"
)

type ConversationDTO struct {
	ConversationId uuid.UUID `json:"conversationId"`
	Name           string    `json:"name"`
	CreatedTime    time.Time `json:"createdTime"`
}

func MapConversationToConversationDTO(conversation *entities.Conversation) *ConversationDTO {
	return &ConversationDTO{
		ConversationId: conversation.Id,
		Name:           conversation.Name,
		CreatedTime:    conversation.CreatedTime,
	}
}
