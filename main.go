package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"resource-service/src/controller"
	"resource-service/src/model"
	"resource-service/utils/constants"
	"resource-service/utils/database"
	applogger "resource-service/utils/logging"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

/*
**********************
Do not edit this file*
**********************
*/

// Create once per go file and re use log
var log *zerolog.Logger = applogger.GetInstance()
var configFilePath *string

func main() {
	configFilePath = flag.String("config-path", "conf/", "conf/")
	flag.Parse()

	loadConfig()

	gin.DisableConsoleColor()
	r := gin.New()
	//Allow CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	setupLogger(r)
	setupRoutes(r)
	setupDatabase()
	startServer(r)
}

// loadConfig - Load the config parameters
func loadConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		if readErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panic().Msgf("No config file found at %s\n", *configFilePath)
		} else {
			log.Panic().Msgf("Error reading config file: %s\n", readErr)
		}
	}
}

// setupDatabase - Set up database
func setupDatabase() {
	database.GetInstance()
	model.Migrate()
}

// startServer - Start server
func startServer(r *gin.Engine) {
	srv := &http.Server{
		Addr:    viper.GetString("server.port"),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("Listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...\n")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("Server forced to shutdown: %s\n", err)
	}

	log.Info().Msg("Server exiting\n")

}

// setupLogger - Configure logging for the server
func setupLogger(r *gin.Engine) {
	// Configure logger
	zlog := applogger.GetInstance()
	r.Use(logger.SetLogger(logger.Config{
		Logger: zlog,
		UTC:    true,
	}))
}

//setupRoutes - Define all the Routes
func setupRoutes(r *gin.Engine) {
	// Instantiate controllers
	resourceController := controller.ResourceController{}
	loginController := controller.LoginController{}
	// Set application context in URL - Do not edit this
	application := r.Group(viper.GetString("server.basepath"))
	{
		// Edit from here
		// All routes for API version V1
		v1 := application.Group("/api/v1")
		{

			resource := v1.Group("/resource")
			{
				resource.POST(constants.LOGIN_USER_PATH, loginController.LoginUser)
				resource.POST(constants.ADD_RESOURCE_PATH, resourceController.AddResource)
				resource.POST(constants.ADD_IMAGE_PATH, resourceController.AddImage)
				resource.POST(constants.ADD_FILE_PATH, resourceController.AddFile)
				resource.POST(constants.GET_FILE_PATH, resourceController.GetFile)
				resource.POST(constants.GET_LIST_OF_RESOURCES_PATH, resourceController.GetListOfResources)
				resource.POST(constants.GET_RESOURCE_PATH, resourceController.GetResource)
				resource.POST(constants.CHANGE_STATUS_PATH, resourceController.ChangeStatus)
				resource.POST(constants.EDIT_RESOURCE_PATH, resourceController.EditResource)
				resource.POST(constants.DELETE_RESOURCE_PATH, resourceController.DeleteResource)
			}
		}
	}
}
