package infrastructure

import (
	"database/sql"
	"fmt"
	"gobrax-api/domain/entities"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"gobrax-api/domain/usecase"
	"gobrax-api/infrastructure/handler"
	"gobrax-api/infrastructure/repository"
)

func Config(env entities.Env, r *gin.Engine, logger *zap.SugaredLogger) (handler *handler.Handler) {
	db := connection(env, logger)
	handler = handlerCfg(db, r, logger)
	return
}

func connection(env entities.Env, logger *zap.SugaredLogger) *sql.DB {
	logger.Info("starting-postgres-connection")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		env.DBHOST, env.DBPORT, env.DBUSER, env.DBPASSAWORD, env.DBNAME)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Error("error message:", zap.Error(err))
		return nil
	}

	err = db.Ping()
	if err != nil {
		logger.Error("error message:", zap.Error(err))
		return nil
	}

	repository.NewPostgres(db, logger)

	return db

}

func handlerCfg(dbConn *sql.DB, r *gin.Engine, logger *zap.SugaredLogger) (handlerStruct *handler.Handler) {
	logger.Info("preparing-modules")

	repositoryTrucker, repositoryVehicle, repositoryAssign := repository.NewPostgres(dbConn, logger)

	ut, uv, ua := usecase.NewUsecase(repositoryTrucker, repositoryVehicle, repositoryAssign, logger)

	return handler.NewHandler(r, ut, uv, ua, logger)

}
