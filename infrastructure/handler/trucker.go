package handler

import (
	"context"
	"encoding/json"
	"errors"
	"gobrax-api/domain/entities"
	errorUtils "gobrax-api/utils/error"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Todo: Swagger
func (h *Handler) createTrucker(c *gin.Context) {
	h.logger.Info("api-handler-create-trucker")

	var (
		ctx     = context.Background()
		payload = Trucker{}
	)

	err := h.getAndCheckTrucker(c, &payload)
	if err != nil {
		h.logger.Info("get-and-check-payload-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(400, errResp)
		return
	}

	err = h.UseCaseTrucker.CreateTrucker(ctx, entities.Trucker{
		Name:        payload.Name,
		Age:         payload.Age,
		Nationality: payload.Nationality,
	})

	if err != nil {
		h.logger.Info("usecase-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	h.logger.Info("trucker-created")
	c.Status(http.StatusOK)

}

func (h *Handler) getTrucker(c *gin.Context) {
	h.logger.Info("api-handler-get-trucker")

	var (
		ctx = context.Background()
	)

	truckers, err := h.UseCaseTrucker.GetTrucker(ctx)
	if err != nil {
		h.logger.Info("usecase-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	truckersResp := []Trucker{}

	for _, x := range truckers {
		truckersResp = append(truckersResp, Trucker{
			ID:          x.ID,
			Name:        x.Name,
			Age:         x.Age,
			Nationality: x.Nationality,
		})
	}

	h.logger.Info("get-truckers-sucessfull")
	c.JSON(http.StatusOK, truckersResp)
}

func (h *Handler) getTruckerByID(c *gin.Context) {
	h.logger.Info("api-handler-create-trucker-by-id")

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

	trucker, err := h.UseCaseTrucker.GetTruckerByID(ctx, convID)
	if err != nil {
		h.logger.Info("usecase-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	h.logger.Info("get-trucker-by-id-sucessfull")
	c.JSON(http.StatusOK, trucker)
}

func (h *Handler) updateTrucker(c *gin.Context) {
	h.logger.Info("api-handler-update-trucker")
	var (
		ctx     = context.Background()
		id      = c.Param("id")
		payload = Trucker{}
	)

	convID, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Info("id-conversion-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	h.getAndCheckTrucker(c, &payload)
	if err != nil {
		h.logger.Info("get-and-check-payload-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(400, errResp)
		return
	}

	err = h.UseCaseTrucker.UpdateTrucker(ctx, entities.Trucker{
		Name:        payload.Name,
		Age:         payload.Age,
		Nationality: payload.Nationality,
	}, convID)
	if err != nil {
		h.logger.Info("usecase-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	h.logger.Info("trucker-updated")
	c.Status(http.StatusOK)
}

func (h *Handler) deleteTrucker(c *gin.Context) {
	h.logger.Info("api-handler-delete-trucker")
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

	err = h.UseCaseTrucker.DeleteTrucker(ctx, convID)
	if err != nil {
		h.logger.Info("usecase-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	h.logger.Info("trucker-deleted")
	c.Status(http.StatusOK)
}

func (h *Handler) getAndCheckTrucker(c *gin.Context, payload *Trucker) (err error) {
	if err = h.unmarshalPayload(c, payload); err != nil {
		return
	}

	return h.checkCreatePayload(payload)
}

func (h *Handler) checkCreatePayload(payload *Trucker) (err error) {
	if payload.Name == "" {
		err = errors.New("invalid name")
		return err
	}

	if payload.Nationality == "" {
		err = errors.New("invalid nationality")
		return err
	}

	if payload.Age >= 0 && payload.Age <= 20 {
		err = errors.New("invalid age")
		return err
	}

	return
}

func (h *Handler) unmarshalPayload(c *gin.Context, payload interface{}) (err error) {
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		return
	}
	return json.Unmarshal(body, payload)
}
