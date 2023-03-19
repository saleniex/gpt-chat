package meta

import "log"

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
	log.Printf("Send mesage: %s", dataStr)
	//TODO perform actual sending
	return nil
}
