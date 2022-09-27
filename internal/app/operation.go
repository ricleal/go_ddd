package app

import (
	"go-cleanarchitecture/internal/app/domain"
	"log"
)

type App struct {
	service      PersonService
	presentation PersonPresentation
}

func New(service PersonService, presentation PersonPresentation) *App {
	return &App{service, presentation}
}

func (a *App) Run() error {
	// Create 10 fake persons
	for i := 0; i < 10; i++ {
		p, err := domain.NewFakePerson()
		if err != nil {
			log.Printf("domain.NewFakePerson() error: %v", err)
			continue
		}
		err = p.Validate()
		if err != nil {
			log.Printf("p.Validate() error: %v", err)
			continue
		}
		err = a.service.Create(*p)
		if err != nil {
			log.Printf("a.service.CreateFakePerson() error: %v", err)
			continue
		}
	}

	p := domain.NewPerson("Invalid Person 123", 111)
	err := p.Validate()
	if err != nil {
		log.Printf("p.Validate() error: %v", err)
	} else {
		err = a.service.Create(*p)
		if err != nil {
			log.Printf("a.service.CreatePerson() error: %v", err)
		}
	}

	persons := a.service.FindAll()
	a.presentation.AddPersons(persons)
	a.presentation.Render()
	return nil
}
