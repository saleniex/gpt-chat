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
	engine.POST("/test", handler.Test{}.Handle)
	engine.GET("/meta/verify-token", handler.VerifyToken{
		Token: env("META_VERIFY_TOKEN"),
	}.Handle)

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
	userLabel := "You"
	aiLabel := "Mobilly"
	prompt := chat.NewPrompt(userLabel, aiLabel)
	promptPath := env("PROMPT_FILE")
	if promptPath != "" {
		if err := prompt.Load(promptPath); err != nil {
			log.Printf("Prompt path %s is not loaded: %s", promptPath, err)
		}
	}
	conversationRepo := chat.NewConversationMemRepo(userLabel, aiLabel, prompt)

	return chat.NewBox(
		env("OPENAI_TOKEN"),
		conversationRepo)
}
