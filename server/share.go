package server

import (
	"encoding/json"
	"net/http"
)

type ShareRequest struct {
	Secret      string `json:"secret"`
	Passphrase  string `json:"passphrase"`
	TTL         uint   `json:"ttl"`
	Recipient   string `json:"recipient"`
	credentials Credentials
}

type ShareResponse struct {
	CustId     string `json:"custid"`
	MetadatKey string `json:"metadata_key"`
	SecretKey  string `json:"secret_key"`
	TTL        uint   `json:"ttl"`
	Updated    uint   `json:"updated"`
	Created    uint   `json:"created"`
}

func Share(shareRequest ShareRequest) (*ShareResponse, error) {
	// body := &ShareRequest{
	// 	Secret:     "olayiwola",
	// 	Passphrase: "ola",
	// 	TTL:        uint(time.Now().UnixNano()) + 120,
	// }
	apiUrl := "https://onetimesecret.com/api/v1/share"

	req, err := http.NewRequest(http.MethodPost, apiUrl, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("secret", shareRequest.Secret)
	q.Add("passphrase", shareRequest.Passphrase)
	q.Add("ttl", string(rune(shareRequest.TTL)))
	req.URL.RawQuery = q.Encode()

	response, err := MakeRequest(req, shareRequest.credentials)
	if err != nil {
		return nil, err
	}

	var shareresponse ShareResponse
	err = json.Unmarshal([]byte(*response), &shareresponse)

	return &shareresponse, err
}
