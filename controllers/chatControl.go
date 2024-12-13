package controllers

import (
	"net/http"
	"strconv"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/gin-gonic/gin"
)

type ChatController struct {
	chatService services.ChatServiceInterface
}

func NewChatController(chatService services.ChatServiceInterface) *ChatController {
	return &ChatController{chatService: chatService}
}

func (ctc *ChatController) ChatController(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "User  ID not found"})
		return
	}

	intUserID, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid User ID type"})
		return
	}

	var input struct {
		UserInput string `json:"user_input"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Input Data Invalid"})
		return
	}

	chat, err := ctc.chatService.ProccessChat(intUserID, input.UserInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": chat})
}

func (ctc *ChatController) GetAllChats(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid User ID"})
		return
	}

	chats, err := ctc.chatService.GetAllChats(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var response []struct {
		Chat entities.Chat `json:"chat"`
	}

	for _, chat := range chats {
		response = append(response, struct {
			Chat entities.Chat `json:"chat"`
		}{
			Chat: chat,
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": response})
}

func (ctc *ChatController) GetChatByID(c *gin.Context) {
	chatIDStr := c.Param("id")
	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid Chat ID"})
		return
	}

	chat, err := ctc.chatService.GetChatByID(chatID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "Chat not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": chat})
}
