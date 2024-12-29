package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
TestPing:
response status code == 200
response body == "OK"
*/
func TestPing(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	// dummy request
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ping(responseRecorder, req)

	res := responseRecorder.Result()

	if res.StatusCode != http.StatusOK {
        t.Errorf("Expected %v, got %v", http.StatusOK, res.StatusCode)
    }

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	if string(body) != "OK" {
        t.Errorf("Expected %v, got %v", "OK", string(body))
    }

}