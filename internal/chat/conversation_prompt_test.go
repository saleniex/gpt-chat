package chat

import (
	"strings"
	"testing"
)

func TestConversationPrompt(t *testing.T) {
	prompt := NewPrompt("USER", "AI")
	if err := prompt.Load("prompt_test.txt"); err != nil {
		t.Errorf("cannot load test prompt file: %s", err)
	}

	repo := NewConversationMemRepo()
	_ = repo.StoreChallengeResponse("_H01_", &ChallengeResponse{
		Challenge: "_CH03_",
		Response:  "_R03_",
	})
	_ = repo.StoreChallengeResponse("_H01_", &ChallengeResponse{
		Challenge: "_CH04_",
		Response:  "_R04_",
	})

	cp := NewConversationPrompt(prompt, repo)
	promptStr := cp.withMessage(&Message{
		Handle: "_H01_",
		Text:   "_MESSAGE_",
	})

	expectedSerializedStr := "_INTRO_TEXT_{br}{br}USER: _CH01_{br}AI: _R01_{br}{br}USER: _CH02_{br}AI: _R02_{br}{br}{br}USER: _CH03_{br}AI: _R03_{br}{br}USER: _CH04_{br}AI: _R04_{br}{br}USER: _MESSAGE_{br}AI: "
	promptSerializedStr := strings.ReplaceAll(promptStr, "\n", "{br}")
	if expectedSerializedStr != promptSerializedStr {
		t.Errorf(
			"prompt string (%d) does not match expected (%d) one",
			len(promptSerializedStr),
			len(expectedSerializedStr))
	}
}
