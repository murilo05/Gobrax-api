package handler

import (
	"gobrax-api/interfaces"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	UseCaseTrucker interfaces.UseCaseTrucker
	UseCaseVehicle interfaces.UseCaseVehicle
	UseCaseAssign  interfaces.UseCaseAssign
	logger         *zap.SugaredLogger
}

func NewHandler(r *gin.Engine, ut interfaces.UseCaseTrucker, uv interfaces.UseCaseVehicle, ua interfaces.RepositoryAssign, logger *zap.SugaredLogger) (handler *Handler) {
	handler = &Handler{
		UseCaseTrucker: ut,
		UseCaseVehicle: uv,
		UseCaseAssign:  ua,
		logger:         logger,
	}

	//Trucker Routes
	RoutesTrucker := r.Group("/trucker")
	{
		RoutesTrucker.POST("/create", handler.createTrucker)
		RoutesTrucker.GET("/", handler.getTrucker)
		RoutesTrucker.GET("/:id", handler.getTruckerByID)
		RoutesTrucker.PUT("/:id", handler.updateTrucker)
		RoutesTrucker.DELETE("/:id", handler.deleteTrucker)
	}

	//Vehicle Routes
	RoutesVehicle := r.Group("/vehicle")
	{
		RoutesVehicle.POST("/create", handler.createVehicle)
		RoutesVehicle.GET("/", handler.getVehicle)
		RoutesVehicle.GET("/:id", handler.getVehicleByID)
		RoutesVehicle.PUT("/:id", handler.updateVehicle)
		RoutesVehicle.DELETE("/:id", handler.deleteVehicle)
	}

	//Assign Routes
	RoutesAssign := r.Group("/assign")
	{
		RoutesAssign.POST("/vehicle/:vehicleID/trucker/:truckerID", handler.assign)
	}

	return nil
}
