package apiv1

import (
	"github.com/ashishthakur913/Floom/api/v1.0/flower"
	"github.com/ashishthakur913/Floom/internal/handlers"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"net/http"
	"io/ioutil"
	"os"
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
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		defer file.Close()

		imageUploadFolder := "../../images";
		if _, err := os.Stat(filepath.Base(imageUploadFolder)); os.IsNotExist(err) {
		    os.Mkdir(filepath.Base(imageUploadFolder), 0700)
		}

		tempFile, err := ioutil.TempFile(filepath.Base(imageUploadFolder), "upload-*.png")
	    if err != nil {
	        fmt.Println(err)
	    }
	    defer tempFile.Close()

	    fileBytes, err := ioutil.ReadAll(file)
	    if err != nil {
	        fmt.Println(err)
	    }
	    tempFile.Write(fileBytes)

	    fmt.Println("Successfully Uploaded File\n")

		c.JSON(http.StatusOK, map[string]string{
			"url": tempFile.Name(),
		})
	})
}
