package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlejandroWaiz/repository-archetype/repository"
	"github.com/AlejandroWaiz/repository-archetype/router/common"
	"github.com/AlejandroWaiz/repository-archetype/router/controller"
	"github.com/AlejandroWaiz/repository-archetype/service"
	"github.com/gorilla/mux"
)

func StartRouter() {

	//port := fmt.Sprintf(":%s", "3000")
	router := mux.NewRouter().StrictSlash(true)

	fmt.Println("Starting...")

	//r := app.CreateRouter()

	router.HandleFunc("/healthcheck", common.HealthCheck).Methods("GET")

	archetypeRepository := repository.NewArchetypeRepository()
	archetypeService := service.NewArchetypeService(archetypeRepository)
	archetypeController := controller.NewArchetypeController(archetypeService)

	router.HandleFunc("/", archetypeController.ArchetypeControllerExample).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))

}
