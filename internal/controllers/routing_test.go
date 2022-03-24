package controllers_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/envelope-zero/backend/internal/test"
	"github.com/stretchr/testify/assert"
)

// TestMain takes care of the test setup for this package
func TestMain(m *testing.M) {
	test.Request("POST", "/v1/budgets", `{ "name": "Morre" }`)
	m.Run()
	os.Remove("data/gorm.db")
}

var getOverviewTests = []struct {
	path     string
	expected string
}{
	{"/", `{ "v1": "/v1" }`},
	{"/v1", `{ "budgets": "/budgets" }`},
}

func TestGetOverview(t *testing.T) {
	for _, tt := range getOverviewTests {
		recorder := test.Request("GET", tt.path, "")

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.JSONEq(t, tt.expected, recorder.Body.String())
	}
}

var optionsHeaderTests = []struct {
	path     string
	expected string
}{
	{"/", "GET"},
	{"/v1", "GET"},
	{"/v1/budgets", "GET, POST"},
	//{"/v1/budgets/1", "GET, PATCH, DELETE"},
}

func TestOptionsHeader(t *testing.T) {
	for _, tt := range optionsHeaderTests {
		recorder := test.Request("OPTIONS", tt.path, "")

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, recorder.Header().Get("allow"), tt.expected)
	}
}
