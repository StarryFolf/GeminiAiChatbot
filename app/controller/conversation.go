package controller

import (
	"fiber/app/helper"
	"fiber/app/models"
	dto "fiber/app/models/dtos"
	"fiber/app/models/entities"
	"fiber/app/models/requests"
	"fiber/pkg/utils"
	"fiber/platform/database"
	"fiber/platform/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

func CreateConversation(ctx *fiber.Ctx) error {
	db := database.GetDB()
	logger.SetUpLogger()
	logr := logger.GetLogger()

	request := &requests.CreateConversationRequest{}
	err := ctx.BodyParser(request)

	if err != nil {
		logr.Errorf("Failed to parse body: %v", err)
		return helper.ResponseWithError(ctx, fiber.StatusBadRequest, "Failed to parse the request!")
	}

	validate := utils.NewValidator()

	if err = validate.Struct(request); err != nil {
		logr.Errorf("Failed to validate fields: %v", err)
		return helper.ResponseWithError(ctx, fiber.StatusBadRequest, "Failed to validate the request!")
	}

	conversationId := uuid.New()
	conversation := &entities.Conversation{
		Id:          conversationId,
		Name:        request.Data.Name,
		CreatedTime: time.Now(),
	}

	err = db.Create(&conversation).Error
	if err != nil {
		logr.Errorf("Failed to create conversation. Error: %v", err)
		return helper.SystemError(ctx)
	}

	for _, sampleMessage := range request.Data.SampleMessages {
		message := &entities.Message{
			Id:             uuid.New(),
			ConversationId: &conversationId,
			Content:        sampleMessage.Data.Content,
			Sender:         sampleMessage.Data.Sender,
			SentTime:       time.Now(),
		}

		err = db.Create(&message).Error

		if err != nil {
			logr.Errorf("Failed to create sample messages for conversation. Error: %v", err)
			return helper.SystemError(ctx)
		}
	}

	conversationDto := dto.MapConversationToConversationDTO(conversation)

	response := models.BaseResponse{
		ResultMessage: "Success!",
		Data:          conversationDto,
	}

	logr.WithField("response", response).Info("Create conversation response")
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func GetAllConversations(ctx *fiber.Ctx) error {
	db := database.GetDB()
	logger.SetUpLogger()
	logr := logger.GetLogger()

	var conversations []*entities.Conversation

	db.Model(conversations).Find(&conversations)

	var conversationDtos []*dto.ConversationDTO

	for _, conversation := range conversations {
		conversationDto := dto.MapConversationToConversationDTO(conversation)
		conversationDtos = append(conversationDtos, conversationDto)
	}

	if conversationDtos == nil {
		conversationDtos = []*dto.ConversationDTO{}
	}

	response := models.BaseResponse{
		ResultMessage: "Success",
		Data:          conversationDtos,
	}

	logr.WithField("response", response).Info("Get all conversations response")
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func GetConversationById(ctx *fiber.Ctx) error {
	db := database.GetDB()
	logger.SetUpLogger()
	logr := logger.GetLogger()

	conversationId, err := uuid.Parse(ctx.Params("conversationId"))
	if err != nil {
		logr.Errorf("Failed to parse conversation Id: %v", err)
		return helper.ResponseWithError(ctx, fiber.StatusBadRequest, "Failed to parse the request!")
	}

	var conversation *entities.Conversation

	err = db.First(&conversation, conversationId).Error
	if err != nil {
		logr.Errorf("Cannot find requested conversation: %v", err)
		return helper.ResponseWithError(ctx, fiber.StatusNotFound, "Requested conversation cannot be found")
	}

	conversationDto := dto.MapConversationToConversationDTO(conversation)

	response := models.BaseResponse{
		ResultMessage: "Success",
		Data:          conversationDto,
	}

	logr.WithField("response", response).Info("Get conversation by id response")
	return ctx.Status(fiber.StatusOK).JSON(response)
}
