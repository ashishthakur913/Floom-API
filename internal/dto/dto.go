package dto

import (
	"mime/multipart"
	"github.com/ashishthakur913/Floom/internal/storage/postgres/models"
)

type GetFlowerResponse struct {
	Id       	int 	`json:"id"`
	ImageURL    string 	`json:"image_url"`
	Name        string 	`json:"name"`
	Price       string 	`json:"price"`
	Rating 		string 	`json:"rating"`
}

type GetFlowersResponse struct {
	Flowers 	[]models.Flower 	`json:"flowers"`
}

type PostFlowerRequest struct {
	ImageURL    string 	`json:"image_url"`
	Name        string 	`json:"name" binding:"required"`
	Price       string 	`json:"price" binding:"required"`
	Rating 		string 	`json:"rating" binding:"required"`
}

type FileUploadRequest struct {
	file    multipart.File 	`json:"file" binding:"required"`
}

type FileUpload struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}