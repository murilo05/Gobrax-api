package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"gobrax-api/domain/entities"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockUseCaseTrucker struct {
	CreateTruckerFunc func(ctx context.Context, trucker entities.Trucker) error
}

func (m *MockUseCaseTrucker) CreateTrucker(ctx context.Context, trucker entities.Trucker) error {
	return nil
}

func (m *MockUseCaseTrucker) GetTrucker(ctx context.Context) ([]entities.Trucker, error) {
	return nil, nil
}

func (m *MockUseCaseTrucker) GetTruckerByID(ctx context.Context, id int) (entities.Trucker, error) {
	return entities.Trucker{}, nil
}

func (m *MockUseCaseTrucker) UpdateTrucker(ctx context.Context, trucker entities.Trucker, id int) error {
	return nil
}

func (m *MockUseCaseTrucker) DeleteTrucker(ctx context.Context, id int) error {
	return nil
}

type ResponseWriterMock struct {
	*httptest.ResponseRecorder
}

func TestCreateTruckerHandler(t *testing.T) {
	h := &Handler{
		UseCaseTrucker: &MockUseCaseTrucker{},
	}

	payload := Trucker{
		Name:        "Murilo",
		Age:         25,
		Nationality: "Brazil",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "http://localhost:8080/trucker/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := &ResponseWriterMock{ResponseRecorder: httptest.NewRecorder()}

	c, _ := gin.CreateTestContext(resp)

	c.Request = req

	h.createTrucker(c)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestCreateTruckerHandler_Error_Empty_Name(t *testing.T) {
	h := &Handler{
		UseCaseTrucker: &MockUseCaseTrucker{},
	}

	payload := Trucker{
		Name:        "",
		Age:         30,
		Nationality: "Brazil",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "http://localhost:8080/trucker/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)
	c.Request = req

	h.createTrucker(c)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestCreateTruckerHandler_Error_Empty_Age(t *testing.T) {
	h := &Handler{
		UseCaseTrucker: &MockUseCaseTrucker{},
	}

	payload := Trucker{
		Name:        "Murilo",
		Age:         0,
		Nationality: "Brazil",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "http://localhost:8080/trucker/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)
	c.Request = req

	h.createTrucker(c)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestCreateTruckerHandler_Error_Empty_Nationality(t *testing.T) {
	h := &Handler{
		UseCaseTrucker: &MockUseCaseTrucker{},
	}

	payload := Trucker{
		Name:        "Murilo",
		Age:         0,
		Nationality: "",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "http://localhost:8080/trucker/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)
	c.Request = req

	h.createTrucker(c)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestCreateTruckerHandler_Age_Under_21(t *testing.T) {
	h := &Handler{
		UseCaseTrucker: &MockUseCaseTrucker{},
	}

	payload := Trucker{
		Name:        "Murilo",
		Age:         18,
		Nationality: "Brazil",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "http://localhost:8080/trucker/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)
	c.Request = req

	h.createTrucker(c)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestGetTruckerHandler(t *testing.T) {
	h := &Handler{
		UseCaseTrucker: &MockUseCaseTrucker{},
	}

	req, _ := http.NewRequest("GET", "http://localhost:8080/trucker/", nil)
	resp := &ResponseWriterMock{ResponseRecorder: httptest.NewRecorder()}

	c, _ := gin.CreateTestContext(resp)

	c.Request = req

	h.getTrucker(c)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetTruckerByIDHandler(t *testing.T) {
	h := &Handler{
		UseCaseTrucker: &MockUseCaseTrucker{},
	}

	req, _ := http.NewRequest("GET", "http://localhost:8080/trucker/1", nil)
	resp := &ResponseWriterMock{ResponseRecorder: httptest.NewRecorder()}

	c, _ := gin.CreateTestContext(resp)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	c.Request = req

	h.getTruckerByID(c)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdateTruckerHandler(t *testing.T) {
	h := &Handler{
		UseCaseTrucker: &MockUseCaseTrucker{},
	}

	payload := Trucker{
		Name:        "Carlos Silva",
		Age:         44,
		Nationality: "Brazil",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", "http://localhost:8080/trucker/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := &ResponseWriterMock{ResponseRecorder: httptest.NewRecorder()}

	c, _ := gin.CreateTestContext(resp)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	c.Request = req

	h.updateTrucker(c)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteTruckerHandler(t *testing.T) {
	h := &Handler{
		UseCaseTrucker: &MockUseCaseTrucker{},
	}

	req, _ := http.NewRequest("DELETE", "http://localhost:8080/trucker/1", nil)
	resp := &ResponseWriterMock{ResponseRecorder: httptest.NewRecorder()}

	c, _ := gin.CreateTestContext(resp)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	c.Request = req

	h.deleteTrucker(c)

	assert.Equal(t, http.StatusOK, resp.Code)
}
