package memory

import "go-cleanarchitecture/internal/app/service"

type PersonRepositoryMem struct {
}

func NewPersonRepositoryMem() *PersonRepositoryMem {
	return &PersonRepositoryMem{}
}

func (r *PersonRepositoryMem) FindAll() []*service.Person {
	return []*service.Person{}
}

func (r *PersonRepositoryMem) Create(p *service.Person) error {
	return nil
}
