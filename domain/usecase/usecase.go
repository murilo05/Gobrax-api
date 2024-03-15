package usecase

import (
	"gobrax-api/interfaces"

	"go.uber.org/zap"
)

type usecase struct {
	repoTrucker interfaces.RepositoryTrucker
	repoVehicle interfaces.RepositoryVehicle
	repoAssign  interfaces.RepositoryAssign
	logger      *zap.SugaredLogger
}

func NewUsecase(repoTrucker interfaces.RepositoryTrucker, repoVehicle interfaces.RepositoryVehicle, repoAssign interfaces.RepositoryAssign, logger *zap.SugaredLogger) (interfaces.UseCaseTrucker, interfaces.RepositoryVehicle, interfaces.RepositoryAssign) {
	return &usecase{
			repoTrucker: repoTrucker,
			logger:      logger,
		}, &usecase{
			repoVehicle: repoVehicle,
			logger:      logger,
		}, &usecase{
			repoAssign: repoAssign,
			logger:     logger,
		}
}
