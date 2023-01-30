package chat

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type ConversationMemRepo struct {
	userLabel string
	aiLabel   string
	history   map[string]string
	prompt    string
}

func NewConversationMemRepo(userLabel, aiLabel string) *ConversationMemRepo {
	return &ConversationMemRepo{
		userLabel: userLabel,
		aiLabel:   aiLabel,
		history:   make(map[string]string),
		prompt:    "The following is conversation between user and assistant.",
	}
}

func (c *ConversationMemRepo) PromptWithMessage(message *Message) string {
	if c.history[message.Handle] == "" {
		c.history[message.Handle] = c.prompt
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

// LoadPrompt loads prompt from file or URL
//
// Content of file should be in form:
// {Intro text}
//
// {USER}: User text
// {AGENT}: Agent response
// ..
func (c *ConversationMemRepo) LoadPrompt(path string) error {
	var content string
	var err error
	if match, _ := regexp.MatchString("^https?://", path); match {
		content, err = contentFromUrl(path)
	} else {
		content, err = contentFromFile(path)
	}

	if err != nil {
		return err
	}

	c.prompt = content
	c.prompt = strings.ReplaceAll(content, "{USER}:", c.userLabel+":")
	c.prompt = strings.ReplaceAll(content, "{AGENT}:", c.aiLabel+":")

	return nil
}

func contentFromFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot load prompt from file %s: %s", filePath, err)
	}
	return string(content), nil
}

func contentFromUrl(path string) (string, error) {
	response, err := http.Get(path)
	if err != nil {
		return "", fmt.Errorf("cannot get from url %s: %s", path, err)
	}
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"cannot get from url %s, server responded with status %d",
			path,
			response.StatusCode)
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("cannot read response from url %s: %s", path, err)
	}

	return string(bytes), nil
}
