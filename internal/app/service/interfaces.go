package service

type PersonRepository interface {
	FindAll() []Person
	Create(p Person) error
}
