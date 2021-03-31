package repository

import (
	"github.com/AlejandroWaiz/repository-archetype/model"
)

type ArchetypeRepository struct {
}

type IArchetypeRepository interface {
	FindAll() model.Archetype
}

func NewArchetypeRepository() *ArchetypeRepository {
	return &ArchetypeRepository{}
}

func (r *ArchetypeRepository) FindAll() model.Archetype {
	archetypeModel := model.Archetype{
		Foo: "bar",
		Bar: "foo",
	}

	return archetypeModel
}
