package handlers

import (
	"github.com/ashishthakur913/Floom/internal/dto"
	"github.com/ashishthakur913/Floom/internal/errors"
	"github.com/ashishthakur913/Floom/internal/storage/postgres"
	"github.com/ashishthakur913/Floom/internal/integrations/s3"
	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	Database     postgres.Database
	S3Uploader   s3.S3
}

type Manager interface {
	GetFlower(ctx *gin.Context) (*dto.GetFlowerResponse, *errors.Error)
	GetFlowers(ctx *gin.Context) (*dto.GetFlowersResponse, *errors.Error)
	SaveFlower(ctx *gin.Context, request dto.PostFlowerRequest) (*dto.GetFlowerResponse, *errors.Error)
	FileUpload(ctx *gin.Context, request *dto.FileUpload) (string, *errors.Error)
}

type manager struct {
	database     postgres.Database
	s3Uploader   s3.S3
}

func NewManager(dependencies *Dependencies) Manager {
	return &manager{
		database:     dependencies.Database,
		s3Uploader:   dependencies.S3Uploader,
	}
}