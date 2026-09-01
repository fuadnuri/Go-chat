package repository

import (
	"github.com/fuadnuri/Go-chat/internal/infrastructure/database/models"
	"gorm.io/gorm"
)

type MessageRepositoryImpl struct {
	DB *gorm.DB
}

func NewMessageRepositoryImpl(db *gorm.DB) *MessageRepositoryImpl {
	return &MessageRepositoryImpl{DB: db}
}

func (r *MessageRepositoryImpl) CreateMessage(message *models.MessageModel) error {
	return r.DB.Create(message).Error
}

func (r *MessageRepositoryImpl) GetMessage(id uint) (*models.MessageModel, error) {
	var message models.MessageModel
	if err := r.DB.Where("id = ?", id).First(&message).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *MessageRepositoryImpl) GetMessagesByConversationID(conversationID uint) ([]models.MessageModel, error) {
	var messages []models.MessageModel
	if err := r.DB.Where("conversation_id = ?", conversationID).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
