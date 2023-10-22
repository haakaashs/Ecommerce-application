package database

import (
	"github.com/haakaashs/antino-labs/config"
	"github.com/haakaashs/antino-labs/models"
	"log"

	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BasicDB interface {
	GetStudents() ([]models.Student, error)
	InsertOrUpdateStudents([]models.Student) error
}

type basicDB struct {
	db *gorm.DB
}

func NewBasicDB() *basicDB {
	return &basicDB{
		db: config.Conn,
	}
}

func (d *basicDB) GetStudents() (students []models.Student, err error) {
	funcDesc := "GetStudents DB"
	log.Println("enter " + funcDesc)

	result := d.db.Debug().Find(&students)
	if err = result.Error; err != nil {
		log.Fatal("error in DB query: ", err.Error())
		return students, err
	}

	log.Println("exit " + funcDesc)
	return students, err
}

func (d *basicDB) InsertOrUpdateStudents(students []models.Student) (err error) {
	funcDesc := "GetStudents DB"
	log.Println("enter " + funcDesc)

	result := d.db.Debug().Save(&students)
	if err = result.Error; err != nil {
		log.Fatal("error in DB query: ", err.Error())
		return err
	}
	log.Println("exit " + funcDesc)
	return nil
}
