package controllers

import (
	"net/http"

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
	var input struct {
		UserID    int    `json:"user_id"`
		UserInput string `json:"user_input"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Input Data Invalid"})
		return
	}
	chat, err := ctc.chatService.ProccessChat(input.UserID, input.UserInput)
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
	c.JSON(http.StatusOK, gin.H{"status": true, "message": chats})
}
