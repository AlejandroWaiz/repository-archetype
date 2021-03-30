package serv

import (
	"encoding/json"
	"net/http"

	"github.com/AlejandroWaiz/repository-archetype/repo"
)

type ArchetypeServ struct {
	Repo repo.IArchetypeRepo
}

type IArchetypeServ interface {
}

func NewArchetypeServ(repo repo.IArchetypeRepo) *ArchetypeServ {
	return &ArchetypeServ{Repo: repo}
}

func (s *ArchetypeServ) ArchetypeServExample() (int, []byte) {
	type Status struct {
		Message string `json:"Message"`
		Details string `json:"Details"`
	}
	status := &Status{Message: "ArchetypeService", Details: "Example"}
	b, _ := json.Marshal(&status)
	return http.StatusOK, b
}
