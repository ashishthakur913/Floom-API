package api

import (
	apiv1 "github.com/ashishthakur913/Floom/api/v1.0"
	"github.com/ashishthakur913/Floom/internal/handlers"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine, manager handlers.Manager) {
	api := r.Group("/api")
	{
		apiv1.ApplyRoutes(api, manager)
	}
}
