package chat

import (
	"fmt"
	"log"
)

type ConversationMemRepo struct {
	userLabel string
	aiLabel   string
	history   map[string]string
	prompt    *Prompt
}

func NewConversationMemRepo(userLabel, aiLabel string, prompt *Prompt) *ConversationMemRepo {
	return &ConversationMemRepo{
		userLabel: userLabel,
		aiLabel:   aiLabel,
		history:   make(map[string]string),
		prompt:    prompt,
	}
}

func (c *ConversationMemRepo) PromptWithMessage(message *Message) string {
	if c.history[message.Handle] == "" {
		c.history[message.Handle] = c.prompt.Content
	}
	newPrompt := fmt.Sprintf(
		"%s\n\n%s: %s\n%s: ",
		c.history[message.Handle],
		c.userLabel,
		message.Text,
		c.aiLabel)
	c.history[message.Handle] = newPrompt

	return newPrompt
}

func (c *ConversationMemRepo) StoreResponse(message *Message) {
	if c.history[message.Handle] == "" {
		log.Fatal("try to store response message while conversation is not opened yet")
	}
	c.history[message.Handle] = c.history[message.Handle] + message.Text
}
