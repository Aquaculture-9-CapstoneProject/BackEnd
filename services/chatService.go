package services

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type ChatServiceInterface interface {
	ProccessChat(userID int, userInput string) (entities.Chat, error)
	GetAllChats(userID int) ([]entities.Chat, error)
	GetChatByID(chatID int) (entities.Chat, error)
}

type chatService struct {
	chatRepo repositories.ChatRepoInterface
}

func NewChatService(chatRepo repositories.ChatRepoInterface) *chatService {
	return &chatService{chatRepo: chatRepo}
}

func (cuc *chatService) ProccessChat(userID int, userInput string) (entities.Chat, error) {
	if userInput == "rekomendasi produk apa ?" || userInput == "rekomendasi produk" {
		products, err := cuc.chatRepo.GetRecommendedProducts()
		if err != nil {
			return entities.Chat{}, err
		}
		aiResponse := "Berikut adalah beberapa rekomendasi produk untuk Anda:\n"
		for _, product := range products {
			aiResponse += fmt.Sprintf("Nama: %s, Kategori: %s, Rating: %.2f\n", product.Nama, product.Kategori, product.Rating)
		}
		chat := entities.Chat{
			UserID:    userID,
			UserInput: userInput,
			AiRespon:  aiResponse,
		}

		savedChat, err := cuc.chatRepo.SaveChat(chat)
		if err != nil {
			return entities.Chat{}, err
		}

		return savedChat, nil
	}

	productName := extractProductName(userInput)
	if productName != "" {
		product, err := cuc.chatRepo.GetProductByName(productName)
		if err != nil {
			return entities.Chat{}, err
		}

		aiResponse := fmt.Sprintf("Berikut adalah detail produk %s:\nNama: %s\nKategori: %s\nHarga: %.2f\nRating: %.2f\nDeskripsi: %s\n",
			product.Nama, product.Nama, product.Kategori, product.Harga, product.Rating, product.Deskripsi)

		chat := entities.Chat{
			UserID:    userID,
			UserInput: userInput,
			AiRespon:  aiResponse,
		}

		savedChat, err := cuc.chatRepo.SaveChat(chat)
		if err != nil {
			return entities.Chat{}, err
		}

		return savedChat, nil
	}

	// Jika bukan pertanyaan produk ljt ke api gemin
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return entities.Chat{}, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(userInput))
	if err != nil {
		return entities.Chat{}, err
	}

	if len(resp.Candidates) == 0 {
		return entities.Chat{}, err
	}

	aiResponse := ""
	for _, candidate := range resp.Candidates {
		if candidate.Content == nil {
			continue
		}
		for _, part := range candidate.Content.Parts {
			aiResponse += fmt.Sprintf("%v", part)
		}
	}

	chat := entities.Chat{
		UserID:    userID,
		UserInput: userInput,
		AiRespon:  aiResponse,
	}

	savedChat, err := cuc.chatRepo.SaveChat(chat)
	if err != nil {
		return entities.Chat{}, err
	}

	return savedChat, nil
}

func (cts *chatService) GetAllChats(userID int) ([]entities.Chat, error) {
	return cts.chatRepo.GetAllChat(userID)
}

func (cts *chatService) GetChatByID(chatID int) (entities.Chat, error) {
	return cts.chatRepo.GetChatByID(chatID)
}

func extractProductName(userInput string) string {
	// mencari nama produk dengan pemrosesan string sederhana atau disbt regex
	re := regexp.MustCompile(`detail produk dari (.*)`)
	matches := re.FindStringSubmatch(strings.ToLower(userInput))

	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
