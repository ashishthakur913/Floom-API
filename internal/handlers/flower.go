package handlers

import (
	"github.com/ashishthakur913/Floom/internal/dto"
	"github.com/ashishthakur913/Floom/internal/errors"
	"github.com/gin-gonic/gin"
	"github.com/ashishthakur913/Floom/internal/storage/postgres/models"
	"strconv"
)

func (m *manager) GetFlower(ctx *gin.Context) (*dto.GetFlowerResponse, *errors.Error) {
	flowerID, _ := strconv.Atoi(ctx.Param("flowerID"))
	flower, err := m.database.GetFlowerTable().GetById(flowerID)
	if err != nil {
		return nil, errors.NotFoundError(err.Error())
	}

	resp := &dto.GetFlowerResponse{
		Id:       	flower.Id,
		ImageURL:	flower.ImageURL,
		Name:		flower.Name,
		Price:		flower.Price,
		Rating:		flower.Rating,
	}
	return resp, nil
}

func (m *manager) GetFlowers(ctx *gin.Context) (*dto.GetFlowersResponse, *errors.Error) {
	flowers, err := m.database.GetFlowerTable().GetAll()
	if err != nil {
		return nil, errors.NotFoundError(err.Error())
	}
	resp := &dto.GetFlowersResponse{
		Flowers:       	flowers,
	}
	return resp, nil
}

func (m *manager) SaveFlower(ctx *gin.Context, request dto.PostFlowerRequest) (*dto.GetFlowerResponse, *errors.Error) {

	/*
		Add Event to DB
	*/
	flower := &models.Flower{
		ImageURL:   request.ImageURL,
		Name:       request.Name,
		Price:      request.Price,
		Rating: 	request.Rating,
	}

	flower, err := m.database.GetFlowerTable().Create(flower)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}

	return &dto.GetFlowerResponse{
		Id:   		flower.Id,
		ImageURL:   flower.ImageURL,
		Name:       flower.Name,
		Price:      flower.Price,
		Rating: 	flower.Rating,
	}, nil
}
