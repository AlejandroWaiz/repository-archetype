package service

import (
	"encoding/json"
	"log"

	"github.com/AlejandroWaiz/repository-archetype/repository"
)

type ArchetypeService struct {
	repository repository.IArchetypeRepository
}

type IArchetypeService interface {
	ArchetypeServiceExample() (error, []byte)
}

func NewArchetypeService(repository repository.IArchetypeRepository) *ArchetypeService {
	return &ArchetypeService{repository: repository}
}

func (s *ArchetypeService) ArchetypeServiceExample() (error, []byte) {
	type Status struct {
		Message string `json:"Message"`
		Details string `json:"Details"`
	}
	status := &Status{Message: "ArchetypeService", Details: "Example"}
	b, err := json.Marshal(&status)

	if err != nil {

		log.Fatal(err)
		return err, nil
	}

	return nil, b
}
