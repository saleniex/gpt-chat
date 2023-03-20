package meta

import (
	"io"
	"net/http"
)

type ServiceResponse struct {
	HttpResponse *http.Response
}

func (r ServiceResponse) Body() string {
	responseBody, readErr := io.ReadAll(r.HttpResponse.Body)
	if readErr != nil {
		return ""
	}
	return string(responseBody)
}

func (r ServiceResponse) IsHttpStatusSuccess() bool {
	return r.HttpResponse.StatusCode >= 200 && r.HttpResponse.StatusCode <= 299
}
