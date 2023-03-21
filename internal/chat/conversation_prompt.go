package chat

import (
	"fmt"
	"strings"
)

// ConversationPrompt Facade to join initial knowledge prompt with conversation repository (context)
type ConversationPrompt struct {
	prompt           *Prompt
	conversationRepo ConversationRepo
}

func NewConversationPrompt(prompt *Prompt, repo ConversationRepo) *ConversationPrompt {
	return &ConversationPrompt{
		prompt:           prompt,
		conversationRepo: repo,
	}
}

func (p ConversationPrompt) withMessage(message *Message) string {
	history := p.conversationRepo.History(message.Handle)
	historyStrBuilder := strings.Builder{}
	for _, historyItem := range history {
		historyStrBuilder.WriteString(fmt.Sprintf(
			"%s: %s\n%s: %s\n\n",
			p.prompt.userLabel,
			historyItem.Challenge,
			p.prompt.aiLabel,
			historyItem.Response))
	}
	return fmt.Sprintf(
		"%s\n\n%s%s: %s\n%s: ",
		p.prompt.Content,
		historyStrBuilder.String(),
		p.prompt.userLabel,
		message.Text,
		p.prompt.aiLabel)
}
