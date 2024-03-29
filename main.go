package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gpt-chat/internal/chat"
	"gpt-chat/internal/handler"
	"gpt-chat/internal/meta"
	"log"
	"os"
)

func main() {
	loadEnvs()

	chatBox := createChatBox()

	engine := gin.Default()
	engine.GET("/", handler.NewRoot().Handle)
	engine.POST("/say", handler.NewSay(chatBox).Handle)
	engine.GET("/meta/webhooks", handler.VerifyToken{
		Token: env("META_VERIFY_TOKEN"),
	}.Handle)
	engine.POST("/meta/webhooks", handler.NewEventNotify(createMetaWebService(), chatBox).Handle)

	if err := engine.Run(env("LISTER_ADDR")); err != nil {
		log.Fatalln("Cannot start web service " + err.Error())
	}
}

func loadEnvs() {
	_ = godotenv.Load(".env.local")
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}
}

func env(key string, defaultVal ...string) string {
	val := os.Getenv(key)
	if val == "" {
		if len(defaultVal) == 0 {
			log.Fatalf("Cannot find environment variable %s value", key)
		}
		return defaultVal[0]
	}
	return val
}

func createChatBox() *chat.Box {
	userLabel := "You"
	aiLabel := "Mobilly"
	prompt := chat.NewPrompt(userLabel, aiLabel)
	promptPath := env("PROMPT_FILE")
	if promptPath != "" {
		if err := prompt.Load(promptPath); err != nil {
			log.Printf("Prompt path %s is not loaded: %s", promptPath, err)
		}
	}
	conversationRepo := chat.NewConversationMemRepo()
	conversationPrompt := chat.NewConversationPrompt(prompt, conversationRepo)

	return chat.NewBox(
		env("OPENAI_TOKEN"),
		conversationRepo,
		conversationPrompt)
}

func createMetaWebService() *meta.Service {
	return &meta.Service{
		FromPhoneNumberId: env("META_FROM_PHONE_NUMBER"),
		AccessToken:       env("META_ACCESS_TOKEN"),
	}
}
