package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type ShareRequest struct {
	Secret      string `json:"secret"`
	Passphrase  string `json:"passphrase"`
	TTL         uint   `json:"ttl"`
	Recipient   string `json:"recipient"`
	Credentials Credentials
}

type ShareResponse struct {
	CustId     string `json:"custid"`
	MetadatKey string `json:"metadata_key"`
	SecretKey  string `json:"secret_key"`
	Link       string
	TTL        uint `json:"ttl"`
	Updated    uint `json:"updated"`
	Created    uint `json:"created"`
}

func Share(shareRequest ShareRequest) (*ShareResponse, error) {
	apiUrl := RootUrl + "share"

	req, err := http.NewRequest(http.MethodPost, apiUrl, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("secret", shareRequest.Secret)
	q.Add("passphrase", shareRequest.Passphrase)
	q.Add("ttl", strconv.Itoa(int(shareRequest.TTL)))
	req.URL.RawQuery = q.Encode()

	response, err := MakeRequest(req, shareRequest.Credentials)
	if err != nil {
		return nil, err
	}

	var shareresponse ShareResponse
	err = json.Unmarshal([]byte(*response), &shareresponse)

	if shareresponse.SecretKey == "" {
		return nil, errors.New("unable to generate secret url, ensure correct parameters are passed")
	}

	shareresponse.Link = SecretUrl + shareresponse.SecretKey

	return &shareresponse, err
}
