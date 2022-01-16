package main

import (
	"bivbonus/internal/handler"
	"bivbonus/internal/repository"
	"bivbonus/internal/service"
	"bivbonus/pkg/config"
	"bivbonus/pkg/middlerware"
	"bivbonus/pkg/security/jwt"
	"bivbonus/pkg/security/jwt/storage"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	log := logrus.New()
	appConfig, err := config.InitConfig(log)
	if err != nil {
		log.Fatal(fmt.Printf("Cannot read appConfig %s", err.Error()))
	}

	database, err := repository.NewDatabase(appConfig.DataBaseConfig)
	if err != nil {
		log.Fatal(fmt.Printf("Cannot open database %s", err.Error()))
	}

	appRepo := repository.NewRepository(database)
	appService := service.NewService(appRepo)

	jwtStorage := storage.NewDatabaseJwtStorage(database)
	security := jwt.NewSecurity(appConfig.JwtConfig, jwtStorage)
	appRouter := handler.NewHandler(appService, security)

	serveAddr := fmt.Sprintf("%s:%s", appConfig.ServerConfig.Host, appConfig.ServerConfig.Port)
	log.Printf("Try to start %s", serveAddr)

	requestHandler := appRouter.InitRouter(middlerware.NewMiddleWare(&jwtStorage, appService)).Handler
	if err := fasthttp.ListenAndServe(serveAddr, requestHandler); err != nil {
		log.Fatal("Cannot start application")
	}
	log.Printf("Server started!")
}
