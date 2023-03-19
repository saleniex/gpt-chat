package meta

import (
	"fmt"
	"io"
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
	httpRequest, newReqErr := http.NewRequest("POST", url, strings.NewReader(dataStr))
	if newReqErr != nil {
		return newReqErr
	}
	httpRequest.Header.Add("Content-Type", "application/json")
	httpRequest.Header.Add("Authorization", "Bearer "+s.AccessToken)
	httpClien := http.Client{}
	response, postErr := httpClien.Do(httpRequest)
	if postErr != nil {
		return postErr
	}
	if response.StatusCode < 200 && response.StatusCode > 299 {
		return fmt.Errorf("server responded with non-success status code %d", response.StatusCode)
	}
	responseBody, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return fmt.Errorf("cannot read server response: %s", readErr)
	}
	log.Printf("Message has been sent %d: %s", response.StatusCode, responseBody)
	return nil
}
