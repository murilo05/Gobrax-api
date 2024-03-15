package handler

import (
	"context"
	"errors"
	"gobrax-api/domain/entities"
	errorUtils "gobrax-api/utils/error"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Todo: Swagger
func (h *Handler) createVehicle(c *gin.Context) {
	h.logger.Info("api-handler-create-vehicle")
	var (
		ctx     = context.Background()
		payload = Vehicle{}
		err     error
	)

	err = h.getAndCheckVehicle(c, &payload)
	if err != nil {
		h.logger.Info("get-and-check-payload-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(400, errResp)
		return
	}

	err = h.UseCaseVehicle.CreateVehicle(ctx, entities.Vehicle{
		Name:  payload.Name,
		Brand: payload.Brand,
		Color: payload.Color,
		Year:  payload.Year,
		Plate: payload.Plate,
	})

	if err != nil {
		h.logger.Info("usecase-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	h.logger.Info("vehicle-created")
	c.Status(http.StatusOK)

}

func (h *Handler) getVehicle(c *gin.Context) {
	h.logger.Info("api-handler-get-vehicle")
	var (
		ctx = context.Background()
	)

	truckers, err := h.UseCaseVehicle.GetVehicle(ctx)
	if err != nil {
		h.logger.Info("usecase-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	vehicleResp := []Vehicle{}

	for _, x := range truckers {
		vehicleResp = append(vehicleResp, Vehicle{
			ID:    x.ID,
			Name:  x.Name,
			Brand: x.Brand,
			Color: x.Color,
			Year:  x.Year,
			Plate: x.Plate,
		})
	}

	h.logger.Info("get-vehicle-sucessfull")
	c.JSON(http.StatusOK, vehicleResp)
}

func (h *Handler) getVehicleByID(c *gin.Context) {
	h.logger.Info("api-handler-get-vehicle-by-id")

	var (
		ctx = context.Background()
		id  = c.Param("id")
	)

	convID, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Info("id-conversion-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	vehicle, err := h.UseCaseVehicle.GetVehicleByID(ctx, convID)
	if err != nil {
		h.logger.Info("usecase-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	vehicleResp := Vehicle{
		ID:    vehicle.ID,
		Name:  vehicle.Name,
		Brand: vehicle.Brand,
		Color: vehicle.Color,
		Year:  vehicle.Year,
		Plate: vehicle.Plate,
	}

	h.logger.Info("get-vehicle-by-id-sucessfull")
	c.JSON(http.StatusOK, vehicleResp)
}

func (h *Handler) updateVehicle(c *gin.Context) {
	h.logger.Info("api-handler-update-vehicle")
	var (
		ctx     = context.Background()
		id      = c.Param("id")
		payload = Vehicle{}
	)

	convID, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Info("id-conversion-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	h.getAndCheckVehicle(c, &payload)
	if err != nil {
		h.logger.Info("get-and-check-payload-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(400, errResp)
		return
	}

	err = h.UseCaseVehicle.UpdateVehicle(ctx, entities.Vehicle{
		Name:  payload.Name,
		Brand: payload.Brand,
		Color: payload.Color,
		Year:  payload.Year,
		Plate: payload.Plate,
	}, convID)
	if err != nil {
		h.logger.Info("usecase-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	h.logger.Info("vehicle-updated")
	c.Status(http.StatusOK)
}

func (h *Handler) deleteVehicle(c *gin.Context) {
	h.logger.Info("api-handler-delete-vehicle")

	var (
		ctx = context.Background()
		id  = c.Param("id")
	)

	convID, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Info("id-conversion-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	err = h.UseCaseVehicle.DeleteVehicle(ctx, convID)
	if err != nil {
		h.logger.Info("usecase-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	h.logger.Info("vehicle-deleted")
	c.Status(http.StatusOK)
}

func (h *Handler) getAndCheckVehicle(c *gin.Context, payload *Vehicle) (err error) {
	if err = h.unmarshalPayload(c, payload); err != nil {
		return
	}

	return h.checkVehiclePayload(payload)
}

func (h *Handler) checkVehiclePayload(payload *Vehicle) (err error) {
	if payload.Name == "" {
		err = errors.New("invalid name")
		return err
	}

	if payload.Brand == "" {
		err = errors.New("invalid brand")
		return err
	}

	if payload.Color == "" {
		err = errors.New("invalid color")
		return err
	}

	if payload.Year == 0 {
		err = errors.New("invalid year")
		return err
	}

	if payload.Plate == "" {
		err = errors.New("invalid plate")
		return err
	}

	return
}
