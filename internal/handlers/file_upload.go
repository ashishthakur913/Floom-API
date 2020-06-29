package handlers

import (
	"fmt"

	"github.com/ashishthakur913/Floom/internal/dto"
	"github.com/ashishthakur913/Floom/internal/errors"
	"github.com/gin-gonic/gin"
)

func (m *manager) FileUpload(ctx *gin.Context, request *dto.FileUpload) (string, *errors.Error) {

	fileName, err := m.s3Uploader.UploadImage(request.File, request.FileHeader)

	if err != nil {
		return "", errors.InternalServerError(fmt.Sprintf("%s", err))
	}
	fmt.Println("File Name: ", fileName)

	return fileName, nil
}
