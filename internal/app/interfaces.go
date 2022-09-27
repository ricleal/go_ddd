package app

import "go-cleanarchitecture/internal/app/domain"

type PersonService interface {
	FindAll() []domain.Person
	Create(p domain.Person) error
}

type PersonPresentation interface {
	Render()
	AddPersons(data []domain.Person)
}
