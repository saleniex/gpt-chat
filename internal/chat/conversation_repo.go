package chat

type ChallengeResponse struct {
	Challenge string
	Response  string
}

// ConversationRepo Conversations consists of history of question (challenge) of one party and response or other
// Repository contain sets of such histories - each history is represented  by conversation handle (identificator)
type ConversationRepo interface {
	// StoreChallengeResponse Store challenge response for conversation handle
	StoreChallengeResponse(handle string, cr *ChallengeResponse) error

	// History Get handle challenge-response history
	History(handle string) []ChallengeResponse
}
