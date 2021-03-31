package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AlejandroWaiz/repository-archetype/serv"
)

type IArchetyepController interface {
}

type ArchetypeController struct {
	Service serv.IArchetypeServ
}

func NewArchetypeController(serv serv.IArchetypeServ) *ArchetypeController {
	return &ArchetypeController{Service: serv}
}

func (controller *ArchetypeController) ArchetypeControllerExample(w http.ResponseWriter, r *http.Request) {
	statusCode, responseBody := controller.Service.ArchetypeServExample()

	if statusCode != nil {
		log.Fatal(statusCode)
		w.WriteHeader(400)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
	w.WriteHeader(200)
}
