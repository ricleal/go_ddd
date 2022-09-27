package presentation

import "go-cleanarchitecture/internal/app/domain"

type PersonPresenter struct {
	presenter PersonPresentation
}

func NewPersonPresenter(presenter PersonPresentation) *PersonPresenter {
	return &PersonPresenter{presenter}
}

func (p *PersonPresenter) AddPersons(data []domain.Person) {
	p.presenter.AddPersons(data)
}

func (p *PersonPresenter) Render() {
	p.presenter.Render()
}
