package handler

import (
	"context"
	errorUtils "gobrax-api/utils/error"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) assign(c *gin.Context) {
	h.logger.Info("api-handler-assign")

	var (
		ctx       = context.Background()
		vehicleID = c.Param("vehicleID")
		truckerID = c.Param("truckerID")
	)

	convVehID, err := strconv.Atoi(vehicleID)
	if err != nil {
		h.logger.Info("id-conversion-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(400, errResp)
		return
	}

	convTruID, err := strconv.Atoi(truckerID)
	if err != nil {
		h.logger.Info("id-conversion-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(400, errResp)
		return
	}

	err = h.UseCaseAssign.Assign(ctx, convVehID, convTruID)

	if err != nil {
		h.logger.Info("usecase-failed")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	h.logger.Info("assign-sucessfull")
	c.Status(http.StatusOK)

}
