package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	expectedResponse := StatusResponse{
		Status: "Nominal",
		Locale: "en",
	}

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			expectedPath := "/status"
			if r.URL.Path != expectedPath {
				t.Errorf("Expected to request %s but got %s instead", expectedPath, r.URL.Path)
			}

			buf := new(bytes.Buffer)
			enc := json.NewEncoder(buf)
			enc.Encode(expectedResponse)

			w.WriteHeader(http.StatusOK)
			w.Write(buf.Bytes())
		}),
	)
	RootUrl = server.URL + "/"
	defer server.Close()

	statusRequest := StatusRequest{
		Credentials: Credentials{
			Username: "test",
			Password: "test",
		},
	}

	response, err := Status(statusRequest)

	if err != nil {
		t.Error(err)
		t.Errorf("Failed to make status request")
	}

	if response.Status != expectedResponse.Status {
		t.Errorf("Expected status %s but got status %s instead", expectedResponse.Status, response.Status)
	}

	if response.Locale != expectedResponse.Locale {
		t.Errorf("Expected locale %s but got locale %s instead", expectedResponse.Locale, response.Locale)
	}

}
