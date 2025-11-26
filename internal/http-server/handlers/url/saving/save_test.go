package saving

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http/httptest"
	"rest-service/internal/http-server/handlers/url/saving/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

type ResponseTest struct {
	Status string `json:"status"`
	Alias  string `json:"alias"`
}

func TestHttpHandler(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	body := bytes.NewReader([]byte(`{"url":"https://google.com"}`))
	req := httptest.NewRequest("POST", "https://google.com", body)
	record := httptest.NewRecorder()

	testURLSaver := mocks.NewURLSaver(t)
	handler := New(log, testURLSaver)
	testURLSaver.On("SaveURL", "https://google.com", mock.Anything).Return(int64(1), nil)
	handler.ServeHTTP(record, req)
	if record.Code != 200 {
		t.Errorf("expected 200, got %d", record.Code)
	}
	respBody := record.Body.String()
	var response ResponseTest
	err := json.Unmarshal([]byte(respBody), &response)
	if err != nil {
		t.Errorf("ailed to unmarshal response: %v", err)

	}
	if response.Status != "OK" {
		t.Errorf("expected status OK,got %s", response.Status)
	}
	if response.Alias == "" {
		t.Errorf("alias must be not empty")
	}
}
