package flower

import (
	"net/http"
	"fmt"
	"log"

	"github.com/ashishthakur913/Floom/internal/common"
	"github.com/ashishthakur913/Floom/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/ashishthakur913/Floom/internal/dto"
)

func ApplyRoutes(v1 *gin.RouterGroup, manager handlers.Manager) {
	flower := v1.Group("/flower")

	
	flower.GET("/:flowerID", func(ctx *gin.Context) {
		resp, err := manager.GetFlower(ctx)
		if err != nil {
			ctx.AbortWithError(err.GetHttpCode(), err.Error())
			return
		}

		ctx.JSON(http.StatusOK, common.ToJSON(resp))
	})
	

	flower.GET("/", func(ctx *gin.Context) {
		resp, err := manager.GetFlowers(ctx)
		if err != nil {
			ctx.AbortWithError(err.GetHttpCode(), err.Error())
			return
		}

		ctx.JSON(http.StatusOK, common.ToJSON(resp))
	})

	flower.POST("/", func(ctx *gin.Context) {
		var request dto.PostFlowerRequest

		if err := ctx.BindJSON(&request); err != nil {
			log.Printf("%v", err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		fmt.Println(request)

		resp, err := manager.SaveFlower(ctx, request)
		if err != nil {
			ctx.AbortWithError(err.GetHttpCode(), err.Error())
			return
		}

		ctx.JSON(http.StatusOK, common.ToJSON(resp))
	})
}
