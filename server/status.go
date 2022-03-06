package server

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
	Locale string `json:"locale"`
}

type StatusRequest struct {
	Credentials Credentials
}

func Status(statusRequest StatusRequest) (*StatusResponse, error) {
	apiUrl := "https://onetimesecret.com/api/v1/status"
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return nil, err
	}

	response, err := MakeRequest(req, statusRequest.Credentials)
	if err != nil {
		return nil, err
	}

	var jsonResponse StatusResponse
	err = json.Unmarshal([]byte(*response), &jsonResponse)
	if err != nil {
		return nil, err
	}

	return &jsonResponse, nil
}
