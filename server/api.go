package server

import (
	"io/ioutil"
	"net/http"
)

type APIResponse string

type Credentials struct {
	Username string
	Password string
}

var RootUrl = "https://onetimesecret.com/api/v1/"

func MakeRequest(req *http.Request, cred Credentials) (*APIResponse, error) {
	client := &http.Client{}

	req.SetBasicAuth(cred.Username, cred.Password)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	response := APIResponse(string(responseBody))

	return &response, nil
}
