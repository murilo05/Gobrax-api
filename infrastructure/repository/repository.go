package repository

import (
	"database/sql"
	"gobrax-api/interfaces"

	"go.uber.org/zap"
)

type mysqlRepo struct {
	DB     *sql.DB
	logger *zap.SugaredLogger
}

func NewPostgres(db *sql.DB, logger *zap.SugaredLogger) (interfaces.RepositoryTrucker, interfaces.RepositoryVehicle, interfaces.RepositoryAssign) {
	return &mysqlRepo{
			DB:     db,
			logger: logger,
		}, &mysqlRepo{
			DB:     db,
			logger: logger,
		}, &mysqlRepo{
			DB:     db,
			logger: logger,
		}
}
