package chat

type ConversationRepo interface {
	// PromptWithMessage get prompt with message added at the end
	PromptWithMessage(message *Message) string

	// StoreResponse store response message at the end of conversations
	StoreResponse(message *Message)
}
