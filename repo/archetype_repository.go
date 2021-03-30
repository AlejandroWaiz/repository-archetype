package repo

import (
	"github.com/AlejandroWaiz/repository-archetype/model"
)

type ArchetypeRepo struct {
}

type IArchetypeRepo interface {
	FindAll() model.Archetype
}

func NewArchetypeRepo() *ArchetypeRepo {
	return &ArchetypeRepo{}
}

func (r *ArchetypeRepo) FindAll() model.Archetype {
	archetypeModel := model.Archetype{
		Foo: "bar",
		Bar: "foo",
	}

	return archetypeModel
}
