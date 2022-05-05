package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestShare(t *testing.T) {
	expectedResponse := ShareResponse{
		CustId:    "Nominal",
		Created:   uint(time.Now().Unix()),
		SecretKey: "REWW",
		TTL:       uint(time.Now().Unix()),
	}

	shareRequest := ShareRequest{
		Secret:     "REE",
		Passphrase: "TEST",
		Credentials: Credentials{
			Username: "test",
			Password: "test",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/share"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected to request %s but got %s instead", expectedPath, r.URL.Path)
		}

		buf := new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.Encode(expectedResponse)

		w.WriteHeader(http.StatusAccepted)
		w.Write(buf.Bytes())
	},
	))

	RootUrl = server.URL + "/"

	response, err := Share(shareRequest)
	if err != nil {
		t.Error(err)
		t.Errorf("Failed to make share request")
	}

	if response.Created != expectedResponse.Created {
		t.Errorf("Expected %v but got %v instead", expectedResponse.Created, response.Created)
	}
}
