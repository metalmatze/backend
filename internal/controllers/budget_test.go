package controllers_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/envelope-zero/backend/internal/models"
	"github.com/envelope-zero/backend/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestNoBudgets(t *testing.T) {
	recorder := test.Request("GET", "/v1/budgets", "")

	var response test.APIListResponse

	err := json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into APIListResponse, '%v'", recorder.Body, err)
	}
	budget := models.Budget{Name: fmt.Sprintf("%v", response.Data[0]["name"])}

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, 1, len(response.Data))
	assert.Equal(t, "Morre", budget.Name)
}

func TestNoBudgetNotFound(t *testing.T) {
	recorder := test.Request("GET", "/v1/budgets/2", "")

	assert.Equal(t, 404, recorder.Code)
}
