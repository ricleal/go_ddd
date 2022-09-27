package service

import (
	"go-cleanarchitecture/internal/app/domain"

	"github.com/google/uuid"
)

type PersonService struct {
	repo PersonRepository
}

func NewPersonService(repo PersonRepository) *PersonService {
	return &PersonService{repo}
}

func (s *PersonService) FindAll() []domain.Person {
	persons := s.repo.FindAll()
	appPersons := make([]domain.Person, len(persons))
	for i, p := range persons {
		appPersons[i] = domain.Person{
			Name: p.Name,
			Age:  p.Age,
		}
	}
	return appPersons
}

func (s *PersonService) Create(p domain.Person) error {
	return s.repo.Create(Person{
		UUID: uuid.Must(uuid.NewRandom()),
		Name: p.Name,
		Age:  p.Age,
	})
}
