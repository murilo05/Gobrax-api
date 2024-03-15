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

type MockUseCaseVehicle struct {
}

func (m *MockUseCaseVehicle) CreateVehicle(ctx context.Context, vehicle entities.Vehicle) error {
	return nil
}

func (m *MockUseCaseVehicle) GetVehicle(ctx context.Context) ([]entities.Vehicle, error) {
	return nil, nil
}

func (m *MockUseCaseVehicle) GetVehicleByID(ctx context.Context, id int) (entities.Vehicle, error) {
	return entities.Vehicle{}, nil
}

func (m *MockUseCaseVehicle) UpdateVehicle(ctx context.Context, vehicle entities.Vehicle, id int) error {
	return nil
}

func (m *MockUseCaseVehicle) DeleteVehicle(ctx context.Context, id int) error {
	return nil
}

func TestCreateVehicleHandler_Success(t *testing.T) {
	h := &Handler{
		UseCaseVehicle: &MockUseCaseVehicle{},
	}

	payload := Vehicle{
		Name:  "Atego",
		Brand: "M.BENZ",
		Color: "White",
		Year:  2021,
		Plate: "ABC1234",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "http://localhost:8080/vehicle/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(resp)
	c.Request = req

	h.createVehicle(c)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestCreateVehicleHandler_Error_Empty_Name(t *testing.T) {
	h := &Handler{
		UseCaseVehicle: &MockUseCaseVehicle{},
	}

	payload := Vehicle{
		Name:  "",
		Brand: "M.BENZ",
		Color: "White",
		Year:  2021,
		Plate: "ABC1234",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "http://localhost:8080/vehicle/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(resp)
	c.Request = req

	h.createVehicle(c)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestCreateVehicleHandler_Error_Empty_Brand(t *testing.T) {
	h := &Handler{
		UseCaseVehicle: &MockUseCaseVehicle{},
	}

	payload := Vehicle{
		Name:  "Atego",
		Brand: "",
		Color: "White",
		Year:  2021,
		Plate: "ABC1234",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "http://localhost:8080/vehicle/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(resp)
	c.Request = req

	h.createVehicle(c)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestCreateVehicleHandler_Error_Empty_Color(t *testing.T) {
	h := &Handler{
		UseCaseVehicle: &MockUseCaseVehicle{},
	}

	payload := Vehicle{
		Name:  "Atego",
		Brand: "M.BENZ",
		Color: "",
		Year:  2021,
		Plate: "ABC1234",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "http://localhost:8080/vehicle/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(resp)
	c.Request = req

	h.createVehicle(c)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestCreateVehicleHandler_Error_Empty_Year(t *testing.T) {
	h := &Handler{
		UseCaseVehicle: &MockUseCaseVehicle{},
	}

	payload := Vehicle{
		Name:  "Atego",
		Brand: "M.BENZ",
		Color: "white",
		Year:  0,
		Plate: "ABC1234",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "http://localhost:8080/vehicle/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(resp)
	c.Request = req

	h.createVehicle(c)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestCreateVehicleHandler_Error_Empty_Plate(t *testing.T) {
	h := &Handler{
		UseCaseVehicle: &MockUseCaseVehicle{},
	}

	payload := Vehicle{
		Name:  "Atego",
		Brand: "M.BENZ",
		Color: "white",
		Year:  2021,
		Plate: "",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "http://localhost:8080/vehicle/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(resp)
	c.Request = req

	h.createVehicle(c)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestGetVehicleHandler_Success(t *testing.T) {
	h := &Handler{
		UseCaseVehicle: &MockUseCaseVehicle{},
	}

	req, _ := http.NewRequest("GET", "http://localhost:8080/vehicle/", nil)
	resp := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(resp)
	c.Request = req

	h.getVehicle(c)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetVehicleByIDHandler_Success(t *testing.T) {
	h := &Handler{
		UseCaseVehicle: &MockUseCaseVehicle{},
	}

	req, _ := http.NewRequest("GET", "http://localhost:8080/vehicle/1", nil)
	resp := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(resp)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	c.Request = req

	h.getVehicleByID(c)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdateVehicleHandler_Success(t *testing.T) {
	h := &Handler{
		UseCaseVehicle: &MockUseCaseVehicle{},
	}

	payload := Vehicle{
		Name:  "Car",
		Brand: "BrandX",
		Color: "Red",
		Year:  2020,
		Plate: "ABC1234",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", "http://localhost:8080/vehicle/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(resp)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	c.Request = req

	h.updateVehicle(c)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteVehicleHandler_Success(t *testing.T) {
	h := &Handler{
		UseCaseVehicle: &MockUseCaseVehicle{},
	}

	req, _ := http.NewRequest("DELETE", "http://localhost:8080/vehicle/1", nil)
	resp := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(resp)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	c.Request = req

	h.deleteVehicle(c)

	assert.Equal(t, http.StatusOK, resp.Code)
}
