package meta

import (
	"fmt"
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

func (s Service) SendRequest(request ServiceRequest) (*ServiceResponse, error) {
	dataStr, dataErr := request.Data()
	if dataErr != nil {
		return nil, dataErr
	}

	httpRequest, newReqErr := s.newHttpPostRequest(dataStr)
	if newReqErr != nil {
		return nil, newReqErr
	}
	httpClient := http.Client{}
	response, postErr := httpClient.Do(httpRequest)
	if postErr != nil {
		return nil, postErr
	}

	return &ServiceResponse{HttpResponse: response}, nil
}

func (s Service) newHttpPostRequest(dataStr string) (*http.Request, error) {
	url := fmt.Sprintf("https://graph.facebook.com/v16.0/%s/messages", s.FromPhoneNumberId)
	httpRequest, newReqErr := http.NewRequest("POST", url, strings.NewReader(dataStr))
	if newReqErr != nil {
		return nil, newReqErr
	}
	httpRequest.Header.Add("Content-Type", "application/json")
	httpRequest.Header.Add("Authorization", "Bearer "+s.AccessToken)

	return httpRequest, nil
}
