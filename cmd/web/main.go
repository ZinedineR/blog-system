package main

import (
	"blog-system/config"
	"blog-system/internal/delivery/http"
	api "blog-system/internal/delivery/http/middleware"
	"blog-system/internal/delivery/http/route"
	"blog-system/internal/repository"
	services "blog-system/internal/services"
	"blog-system/migration"
	"blog-system/pkg/database"
	"blog-system/pkg/httpclient"
	"blog-system/pkg/logger"
	"blog-system/pkg/server"
	"blog-system/pkg/signature"
	"blog-system/pkg/xvalidator"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var (
	httpClient    httpclient.Client
	sqlClientRepo *database.Database
)

// @title           Pigeon
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/notificationsvc/api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	validate, _ := xvalidator.NewValidator()
	conf := config.InitAppConfig(validate)
	logger.SetupLogger(&logger.Config{
		AppENV:  conf.AppEnvConfig.AppEnv,
		LogPath: conf.AppEnvConfig.LogFilePath,
		Debug:   conf.AppEnvConfig.AppDebug,
	})
	initInfrastructure(conf)
	ginServer := server.NewGinServer(&server.GinConfig{
		HttpPort:     conf.AppEnvConfig.HttpPort,
		AllowOrigins: conf.AppEnvConfig.AllowOrigins,
		AllowMethods: conf.AppEnvConfig.AllowMethods,
		AllowHeaders: conf.AppEnvConfig.AllowHeaders,
	})
	signaturerService := signature.NewSignature(conf.AppEnvConfig.JWTSecret, "")
	// repository
	userRepository := repository.NewUserSQLRepository()
	postRepository := repository.NewPostSQLRepository()
	commentRepository := repository.NewCommentSQLRepository()

	// service
	userService := services.NewUserService(sqlClientRepo.GetDB(), userRepository, signaturerService, validate)
	postService := services.NewPostService(sqlClientRepo.GetDB(), postRepository, userRepository, signaturerService, validate)
	commentService := services.NewCommentService(sqlClientRepo.GetDB(), commentRepository, postRepository, signaturerService, validate)
	// Handler
	userHandler := http.NewUserHTTPHandler(userService)
	postHandler := http.NewPostHTTPHandler(postService)
	commentHandler := http.NewCommentHTTPHandler(commentService)

	router := route.Router{
		App:            ginServer.App,
		Middleware:     api.NewMiddleware(signaturerService),
		UserHandler:    userHandler,
		PostHandler:    postHandler,
		CommentHandler: commentHandler,
	}
	router.Setup()
	//router.SwaggerRouter()
	echan := make(chan error)
	go func() {
		echan <- ginServer.Start()
	}()

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	select {
	case <-term:
		slog.Info("signal terminated detected")
	case err := <-echan:
		slog.Error("Failed to start http server", err)
	}
}

func initInfrastructure(config *config.Config) {
	sqlClientRepo = initSQL(config)
	httpClient = initHttpclient()
}

func initSQL(conf *config.Config) *database.Database {
	db := database.NewDatabase(conf.DatabaseConfig.Dbservice, &database.Config{
		DbHost:   conf.DatabaseConfig.Dbhost,
		DbUser:   conf.DatabaseConfig.Dbuser,
		DbPass:   conf.DatabaseConfig.Dbpassword,
		DbName:   conf.DatabaseConfig.Dbname,
		DbPort:   strconv.Itoa(conf.DatabaseConfig.Dbport),
		DbPrefix: conf.DatabaseConfig.DbPrefix,
	})
	if conf.IsStaging() {
		migration.AutoMigration(db)
	}
	return db
}

func initHttpclient() httpclient.Client {
	httpClientFactory := httpclient.New()
	httpClient := httpClientFactory.CreateClient()
	return httpClient
}
