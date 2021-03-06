package server

import (
	"encoding/json"
	"net/http"
)

type RetrieveRequest struct {
	Secret      string `json:"SECRET_KEY"`
	Passphrase  string `json:"passphrase"`
	Credentials Credentials
}

type RetrieveResponse struct {
	Value  string `json:"value"`
	Secret string `json:"secret_key"`
}

func Retrieve(retrieveRequest RetrieveRequest) (*RetrieveResponse, error) {
	apiUrl := RootUrl + "secret/" + retrieveRequest.Secret

	req, err := http.NewRequest(http.MethodPost, apiUrl, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("passphrase", retrieveRequest.Passphrase)
	req.URL.RawQuery = q.Encode()

	response, err := MakeRequest(req, retrieveRequest.Credentials)
	if err != nil {
		return nil, err
	}

	var retrieveResponse RetrieveResponse
	err = json.Unmarshal([]byte(*response), &retrieveResponse)

	return &retrieveResponse, err
}
