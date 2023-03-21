package chat

// ConversationMemRepo ConversationRepo implementation where data is stored in-memory
// exists only while application is active
type ConversationMemRepo struct {
	// MaxHandleCount Max items per handle, when reached remove older
	MaxHandleCount int
	history        map[string][]ChallengeResponse
}

func NewConversationMemRepo() *ConversationMemRepo {
	return &ConversationMemRepo{
		history:        make(map[string][]ChallengeResponse),
		MaxHandleCount: 10,
	}
}

func (c *ConversationMemRepo) StoreChallengeResponse(handle string, cr *ChallengeResponse) error {
	if c.history[handle] == nil {
		c.history[handle] = make([]ChallengeResponse, 0)
	}

	if len(c.history[handle]) >= c.MaxHandleCount {
		c.history[handle] = c.history[handle][1:]
	}

	c.history[handle] = append(c.history[handle], *cr)

	return nil
}

func (c *ConversationMemRepo) History(handle string) []ChallengeResponse {
	return c.history[handle]
}
