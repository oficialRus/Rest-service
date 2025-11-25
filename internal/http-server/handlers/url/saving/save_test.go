package saving

import (
	"bytes"
	"io"
	"log/slog"
	"net/http/httptest"
	"rest-service/internal/http-server/handlers/url/saving/mocks"
	"testing"
)

func TestHttpHandler(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	body := bytes.NewReader([]byte(`{"url":"https://google.com"}`))
	req := httptest.NewRequest("POST", "https://google.com", body)
	record := httptest.NewRecorder()

	testURLSaver := mocks.NewURLSaver(t)
	handler := New(log, testURLSaver)
	handler.ServeHTTP(record, req)

}
