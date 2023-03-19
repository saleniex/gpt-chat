package meta

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Service struct {
	FromPhoneNumberId string
	AccessToken       string
}

type ServiceRequest interface {
	Data() (string, error)
}

func (s Service) SendRequest(request ServiceRequest) error {
	dataStr, err := request.Data()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("https://graph.facebook.com/v16.0/%s/messages", s.FromPhoneNumberId)
	log.Printf("Send mesage: %s", dataStr)

	response, postErr := http.Post(url, "application/json", strings.NewReader(dataStr))
	if postErr != nil {
		return postErr
	}
	if response.StatusCode < 200 && response.StatusCode > 299 {
		return fmt.Errorf("server responded with non-success status code %d", response.StatusCode)
	}
	return nil
}
