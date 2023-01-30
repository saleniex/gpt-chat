package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gpt-chat/internal/chat"
	"gpt-chat/internal/handler"
	"log"
	"os"
)

func main() {
	loadEnvs()

	engine := gin.Default()
	engine.GET("/", handler.NewRoot().Handle)
	engine.POST("/say", handler.NewSay(chatBox()).Handle)

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

func chatBox() *chat.Box {
	conversationRepo := chat.NewConversationMemRepo("You", "Mobilly")
	promptFilePath := env("PROMPT_FILE")
	if promptFilePath != "" {
		if err := conversationRepo.LoadPrompt(promptFilePath); err != nil {
			log.Printf("Prompt file %s is not loaded: %s", promptFilePath, err)
		}
	}

	return chat.NewBox(
		env("OPENAI_TOKEN"),
		conversationRepo)
}
