package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRetrieve(t *testing.T) {
	expectedResponse := RetrieveResponse{
		Value:  "test",
		Secret: "test",
	}
	retrieveRequest := RetrieveRequest{
		Secret:     "test",
		Passphrase: "test",
		Credentials: Credentials{
			Username: "test",
			Password: "test",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := fmt.Sprintf("/secret/%s", retrieveRequest.Secret)
		if r.URL.Path != expectedPath {
			t.Errorf("Expected to request %s but got %s instead", expectedPath, r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		buf := new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.Encode(expectedResponse)

		w.Write(buf.Bytes())
	}))

	RootUrl = server.URL + "/"
	response, err := Retrieve(retrieveRequest)

	if err != nil {
		t.Error(err)
		t.Errorf("Failed to make retrieve request")
	}

	if response.Secret != expectedResponse.Secret {
		t.Errorf("Expected %v but got %v instead", expectedResponse.Secret, response.Secret)
	}

	if response.Value != expectedResponse.Value {
		t.Errorf("Expected %v but got %v instead", expectedResponse.Value, response.Value)
	}
}
