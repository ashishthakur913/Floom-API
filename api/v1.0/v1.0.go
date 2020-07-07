package apiv1

import (
	"github.com/ashishthakur913/Floom/api/v1.0/flower"
	"github.com/ashishthakur913/Floom/internal/handlers"
	"github.com/ashishthakur913/Floom/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup, manager handlers.Manager) {
	v1 := r.Group("/v1.0")
	{
		flower.ApplyRoutes(v1, manager)
		fileUpload(v1, manager)
	}
}

func fileUpload(v1 *gin.RouterGroup, manager handlers.Manager) {
	fileUpload := v1.Group("/file")
	fileUpload.POST("/", func(c *gin.Context) {
		file, fileHeader, err := c.Request.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		fileUploader := &dto.FileUpload{
			File:       file,
			FileHeader: fileHeader,
		}
		fmt.Sprintf("\nFile :%s\n", fileHeader.Filename)

		fileName, err2 := manager.FileUpload(c, fileUploader)
		if err2 != nil {
			fmt.Println("Err", err2)
			c.JSON(err2.GetHttpCode(), map[string]interface{}{
				"error": err2.GetMessage(),
			})
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"url": fileName,
		})
	})
}
