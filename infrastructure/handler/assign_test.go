package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockUseCaseAssign struct {
	AssignFunc func(ctx context.Context, vehicleID, truckerID int) error
}

func (m *MockUseCaseAssign) Assign(ctx context.Context, vehicleID, truckerID int) error {
	return nil
}

func TestAssignHandler_Success(t *testing.T) {
	h := &Handler{
		UseCaseAssign: &MockUseCaseAssign{},
	}

	req, _ := http.NewRequest("POST", "http://localhost:8080/assign/123/456", nil)
	resp := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(resp)
	c.Request = req

	c.Params = []gin.Param{
		{Key: "vehicleID", Value: "123"},
		{Key: "truckerID", Value: "456"},
	}

	h.assign(c)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestAssignHandler_Error_Invalid_VehicleID(t *testing.T) {
	h := &Handler{
		UseCaseAssign: &MockUseCaseAssign{},
	}

	req, _ := http.NewRequest("POST", "http://localhost:8080/assign/123/456", nil)
	resp := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(resp)
	c.Request = req

	c.Params = []gin.Param{
		{Key: "vehicleID", Value: "ABC"},
		{Key: "truckerID", Value: "456"},
	}

	h.assign(c)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestAssignHandler_Error_Invalid_TruckerID(t *testing.T) {
	h := &Handler{
		UseCaseAssign: &MockUseCaseAssign{},
	}

	req, _ := http.NewRequest("POST", "http://localhost:8080/assign/123/456", nil)
	resp := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(resp)
	c.Request = req

	c.Params = []gin.Param{
		{Key: "vehicleID", Value: "123"},
		{Key: "truckerID", Value: "ABC"},
	}

	h.assign(c)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestMockUseCaseAssign_Assign(t *testing.T) {
	type args struct {
		ctx       context.Context
		vehicleID int
		truckerID int
	}
	tests := []struct {
		name    string
		m       *MockUseCaseAssign
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Assign(tt.args.ctx, tt.args.vehicleID, tt.args.truckerID); (err != nil) != tt.wantErr {
				t.Errorf("MockUseCaseAssign.Assign() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
