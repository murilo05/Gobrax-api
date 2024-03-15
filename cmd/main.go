package main

import (
	"gobrax-api/domain/entities"
	"gobrax-api/infrastructure"
	"gobrax-api/utils/env"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := gin.Default()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	log := logger.Sugar()
	defer logger.Sync()

	HOST := env.GetString("HOST")
	PORT := env.GetString("PORT")
	DBHOST := env.GetString("DBHOST")
	DBPORT := env.GetInt("DBPORT")
	DBUSER := env.GetString("DBUSER")
	DBPASSAWORD := env.GetString("DBPASSWORD")
	DBNAME := env.GetString("DBNAME")

	env := entities.Env{
		HOST:        HOST,
		PORT:        PORT,
		DBHOST:      DBHOST,
		DBPORT:      DBPORT,
		DBUSER:      DBUSER,
		DBPASSAWORD: DBPASSAWORD,
		DBNAME:      DBNAME,
	}

	infrastructure.Config(env, r, log)

	r.Run(HOST + ":" + PORT)
}
