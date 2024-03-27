package entities

import (
	"github.com/google/uuid"
	"time"
)

type Conversation struct {
	Id          uuid.UUID `gorm:"type:uuid;primaryKey;" json:"conversationId"`
	Name        string    `gorm:"type:varchar(100);not null;" json:"name"`
	CreatedTime time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"createdTime"`

	Messages []Message `gorm:"foreignKey:ConversationId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (Conversation) TableName() string {
	return "Conversations"
}
