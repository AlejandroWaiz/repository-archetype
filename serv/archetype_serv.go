package serv

import (
	"encoding/json"

	"github.com/AlejandroWaiz/repository-archetype/repo"
)

type ArchetypeServ struct {
	Repo repo.IArchetypeRepo
}

type IArchetypeServ interface {
	ArchetypeServExample() (error, []byte)
}

func NewArchetypeServ(repo repo.IArchetypeRepo) *ArchetypeServ {
	return &ArchetypeServ{Repo: repo}
}

func (s *ArchetypeServ) ArchetypeServExample() (error, []byte) {
	type Status struct {
		Message string `json:"Message"`
		Details string `json:"Details"`
	}
	status := &Status{Message: "ArchetypeService", Details: "Example"}
	b, err := json.Marshal(&status)

	return err, b
}
