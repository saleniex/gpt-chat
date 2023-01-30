package chat

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Prompt struct {
	content   string
	userLabel string
	aiLabel   string
}

func NewPrompt() *Prompt {
	return &Prompt{}
}

func (p *Prompt) Load(path string) error {
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

	p.content = content
	p.content = strings.ReplaceAll(content, "{USER}:", p.userLabel+":")
	p.content = strings.ReplaceAll(content, "{AGENT}:", p.aiLabel+":")

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
