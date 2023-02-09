package domain

import (
	"fmt"
	"regexp"

	"github.com/go-faker/faker/v4"
)

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func NewFakePerson() (*Person, error) {

	f := struct {
		FirstName string `faker:"first_name"`
		LastName  string `faker:"last_name"`
		Age       int    `faker:"boundary_start=21, boundary_end=112"`
	}{}

	err := faker.FakeData(&f)
	if err != nil {
		return nil, err
	}

	return &Person{
		Name: fmt.Sprintf("%s %s", f.FirstName, f.LastName),
		Age:  f.Age,
	}, nil
}

func (s *Person) String() string {
	return fmt.Sprintf("%s (%d)", s.Name, s.Age)
}

func (s *Person) Equals(p Person) bool {
	return s.Name == p.Name && s.Age == p.Age
}

func (s *Person) EqualsPtr(p *Person) bool {
	return s.Name == p.Name && s.Age == p.Age
}

func (s *Person) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("name is empty")
	}

	numeric := regexp.MustCompile(`\d`).MatchString(s.Name)
	if numeric {
		return fmt.Errorf("name contains numeric characters")
	}

	if s.Age < 0 {
		return fmt.Errorf("age is negative")
	}
	return nil
}
