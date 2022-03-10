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

var url = "https://onetimesecret.com/"
var RootUrl = url + "api/v1/"
var SecretUrl = url + "secret/"

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
