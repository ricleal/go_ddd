package presentation

import "go-cleanarchitecture/internal/app/domain"

type PersonPresentation interface {
	Render()
	AddPersons(data []domain.Person)
}
