package entities

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Id       uuid.UUID `gorm:"type:uuid;primaryKey;" json:"messageId"`
	Content  string    `gorm:"type:varchar(3000);not null;" json:"content"`
	Sender   string    `gorm:"type:varchar(20);not null;" json:"sender"`
	SentTime time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"sentTime"`

	ConversationId *uuid.UUID `gorm:"type:uuid;not null;" json:"conversationId"`
}

func (Message) TableName() string {
	return "Messages"
}
