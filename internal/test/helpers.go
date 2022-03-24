package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/envelope-zero/backend/internal/controllers"
)

// Request is a helper method to simplify making a HTTP request for tests
func Request(method string, url string, body string) httptest.ResponseRecorder {
	byteStr := []byte(body)

	router := controllers.Router()

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(byteStr))
	router.ServeHTTP(recorder, req)

	return *recorder
}

type APIListResponse struct {
	Data  []map[string]interface{}
	Links []map[string]string
	Error string
}
