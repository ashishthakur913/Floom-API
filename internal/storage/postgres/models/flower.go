package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Flower struct {
	Id       	int
	ImageURL    string
	Name        string
	Price       string
	Rating 		string
}

type flowerTable struct {
	db *gorm.DB
}

func NewFlowerTable(db *gorm.DB) FlowerTable {
	return &flowerTable{
		db: db,
	}
}

type FlowerTable interface {
	Create(flower *Flower) (*Flower, error)
	GetById(id int) (*Flower, error)
	GetAll() ([]Flower, error)
}

func (u *flowerTable) Create(flower *Flower) (*Flower, error) {
	if err := u.db.Create(&flower).Scan(&flower).Error; err != nil {
		return nil, err
	}
	return flower, nil
}

func (u *flowerTable) GetById(id int) (*Flower, error) {
	flower := &Flower{}
	query := fmt.Sprintf("id=%d", id)
	if err := u.db.Where(query).First(&flower).Error; err != nil {
		return nil, err
	}
	return flower, nil
}

func (u *flowerTable) GetAll() ([]Flower, error) {
 	flowers := []Flower{}
	if err := u.db.Find(&flowers).Error; err != nil {
		return nil, err
	}
	return flowers, nil
}
