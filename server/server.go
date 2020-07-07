package server

import (
	"github.com/ashishthakur913/Floom/api"
	limits "github.com/gin-contrib/size"
	"github.com/ashishthakur913/Floom/config"
	"github.com/ashishthakur913/Floom/internal/handlers"
	"github.com/ashishthakur913/Floom/internal/storage/postgres"
	"github.com/ashishthakur913/Floom/internal/integrations/s3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

// To be called from application.go to initialize server with required dependencies
func Run() {
	config := config.Initialize()
	// initializes database

	database, err := postgres.Initialize(config)
	if err != nil {
		panic(err.Error())
	}

	s3Uploader, err := s3.Initialize()
	if err != nil {
		panic(err.Error())
	}

	dependencies := &handlers.Dependencies{
		Database:     database,
		S3Uploader:   s3Uploader,
	}
	manager := handlers.NewManager(dependencies)

	app := gin.Default() // create gin app

	app.Use(limits.RequestSizeLimiter(2500000))

	app.Use(cors.Default())
	app.Static("/images", filepath.Base("../images"))

	gin.SetMode(config.Server().Mode())

	api.ApplyRoutes(app, manager) // apply api router
	err = app.Run(":" + config.Server().Port())
	panic(err.Error())
}
