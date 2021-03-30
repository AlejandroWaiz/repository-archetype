package controller

import (
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

func (controller *ArchetypeController) ArchetypeControllerExample()
