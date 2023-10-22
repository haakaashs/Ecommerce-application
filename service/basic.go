package service

import (
	"github.com/haakaashs/antino-labs/database"
	"github.com/haakaashs/antino-labs/models"
	"log"
	// "github.com/gin-gonic/gin"
)

type BasicService interface {
	GetStudents() ([]models.Student, error)
	InsertOrUpdateStudents([]models.Student) error
}

type basicService struct {
	basicDB database.BasicDB
}

func NewBasicService(basicDB database.BasicDB) *basicService {
	return &basicService{
		basicDB: basicDB,
	}
}

func (s *basicService) GetStudents() ([]models.Student, error) {
	funcDesc := "GetStudents Service"
	log.Println("enter " + funcDesc)

	students, err := s.basicDB.GetStudents()
	if err != nil {
		log.Fatal(err.Error())
		return students, err
	}
	log.Println("exit " + funcDesc)
	return students, err
}

func (s *basicService) InsertOrUpdateStudents(students []models.Student) (err error) {
	funcDesc := "GetStudents service"
	log.Println("enter " + funcDesc)

	err = s.basicDB.InsertOrUpdateStudents(students)
	if err != nil {
		log.Fatal("error in service query: ", err.Error())
		return err
	}
	log.Println("exit " + funcDesc)
	return nil
}
