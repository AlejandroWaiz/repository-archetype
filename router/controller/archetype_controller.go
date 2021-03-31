package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AlejandroWaiz/repository-archetype/service"
)

type IArchetyepController interface {
	ArchetypeControllerExample(w http.ResponseWriter, r *http.Request)
}

type ArchetypeController struct {
	Service service.IArchetypeService
}

func NewArchetypeController(service service.IArchetypeService) *ArchetypeController {
	return &ArchetypeController{Service: service}
}

func (controller *ArchetypeController) ArchetypeControllerExample(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here is")
	statusCode, responseBody := controller.Service.ArchetypeServiceExample()

	if statusCode != nil {
		log.Fatal(statusCode)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
	w.WriteHeader(http.StatusOK)
}
