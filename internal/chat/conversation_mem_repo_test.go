package chat

import (
	"testing"
)

func TestConversationMemRepo(t *testing.T) {
	r := NewConversationMemRepo()
	r.MaxHandleCount = 2

	_ = r.StoreChallengeResponse("_H1_", &ChallengeResponse{
		Challenge: "_CH1_",
		Response:  "_RE1_",
	})
	_ = r.StoreChallengeResponse("_H1_", &ChallengeResponse{
		Challenge: "_CH2_",
		Response:  "_RE2_",
	})
	_ = r.StoreChallengeResponse("_H2_", &ChallengeResponse{
		Challenge: "_CH3_",
		Response:  "_RE3_",
	})

	h1 := r.History("_H1_")
	if len(h1) != 2 {
		t.Errorf("Handle history H1 len mismatch")
	}
	h2 := r.History("_H2_")
	if len(h2) != 1 {
		t.Errorf("Handle history H2 len mismatch")
	}

	_ = r.StoreChallengeResponse("_H1_", &ChallengeResponse{
		Challenge: "_CH3_",
		Response:  "_RE3_",
	})
	h3 := r.History("_H1_")
	if len(h3) != 2 {
		t.Errorf("Max handle count exceeded")
	}
}
