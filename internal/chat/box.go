package chat

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	gogpt "github.com/sashabaranov/go-gpt3"
	"strings"
)

type Box struct {
	accessToken        string
	client             *gogpt.Client
	repo               ConversationRepo
	conversationPrompt *ConversationPrompt
}

func NewBox(accessToken string, repo ConversationRepo, conversationPrompt *ConversationPrompt) *Box {
	return &Box{
		accessToken:        accessToken,
		repo:               repo,
		conversationPrompt: conversationPrompt,
	}
}

func (b *Box) ResponseOn(message *Message) (*Message, error) {
	if message.Handle == "" {
		message.Handle = uuid.New().String()
	}
	request := b.messageRequest(message)

	response, err := b.gptClient().CreateCompletion(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("error while creating completition: %s", err)
	}

	responseMessageText := strings.TrimSpace(response.Choices[0].Text)

	repoErr := b.repo.StoreChallengeResponse(message.Handle, &ChallengeResponse{
		Challenge: message.Text,
		Response:  responseMessageText,
	})
	if repoErr != nil {
		return nil, repoErr
	}

	return &Message{
		Handle: message.Handle,
		Text:   responseMessageText,
	}, nil
}

func (b *Box) gptClient() *gogpt.Client {
	if b.client == nil {
		b.client = gogpt.NewClient(b.accessToken)
	}

	return b.client
}

func (b *Box) messageRequest(message *Message) gogpt.CompletionRequest {
	return gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci002,
		Temperature: 0.7,
		MaxTokens:   400,
		TopP:        1.0,
		Prompt:      b.conversationPrompt.withMessage(message),
		Stop:        []string{},
	}
}
