package postgres

import (
	"fmt"
	"log"

	"github.com/ashishthakur913/Floom/config"
	"github.com/ashishthakur913/Floom/internal/storage/postgres/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database interface {
	GetFlowerTable() models.FlowerTable
}

type database struct {
	db *gorm.DB
	flowerTable             models.FlowerTable
}

// Initialize initializes the database
func Initialize(config config.Config) (*database, error) {
	user := config.Database().User()
	password := config.Database().Password()
	databaseName := config.Database().Name()
	Host := config.Database().Host()
	port := config.Database().Port()

	var err error

	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", Host, port, user, databaseName, password)

	db, err := gorm.Open("postgres", connectString)
	db.LogMode(true) // logs SQL
	if err != nil {
		panic(err)
	}

	log.Printf("Connected to database")
	models.Migrate(db)

	dbObject := &database{
		db: db,

		//Initializing individual table object
		flowerTable:             models.NewFlowerTable(db),
	}

	return dbObject, err
}

/*
	Returns users table interface
*/
func (d *database) GetFlowerTable() models.FlowerTable {
	return d.flowerTable
}
