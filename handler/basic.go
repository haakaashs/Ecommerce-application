package handler

import (
	"log"
	"net/http"

	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/service"
)

type BasicHandler interface {
	GetStudents(http.ResponseWriter, *http.Request)
	InsertOrUpdateStudents(http.ResponseWriter, *http.Request)
}

type basicHandler struct {
	basicService service.BasicService
}

func NewBasicHandler(basicService service.BasicService) *basicHandler {
	return &basicHandler{
		basicService: basicService,
	}
}

func (h *basicHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	funcDesc := "GetStudents Handler"
	log.Println("enter " + funcDesc)

	_, err := h.basicService.GetStudents()
	if err != nil {
		log.Fatal("Unable to get students: ", err.Error())
	}

	log.Println("exit " + funcDesc)
	// ctx.JSON(http.StatusOK, students)
}

func (h *basicHandler) InsertOrUpdateStudents(w http.ResponseWriter, r *http.Request) {
	funcDesc := "GetStudents Handler"
	log.Println("enter " + funcDesc)

	students := []models.Student{}
	// ctx.BindJSON(&students)

	err := h.basicService.InsertOrUpdateStudents(students)
	if err != nil {
		log.Fatal("error in "+funcDesc, err.Error())
	}

	log.Println("exit " + funcDesc)
	// ctx.JSON(http.StatusOK, `{message:operation successful}`)
}
