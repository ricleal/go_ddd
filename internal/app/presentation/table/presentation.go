package table

import (
	"go-cleanarchitecture/internal/app/domain"
	"strconv"
)

type Table struct {
	table
}

func NewTable() *Table {
	return &Table{
		table: table{
			headers: []string{"Name", "Age"},
		},
	}
}

func (t *Table) Render() {
	t.render()
}

func (t *Table) AddPersons(data []domain.Person) {
	for _, person := range data {
		t.addRow([]string{person.Name, strconv.Itoa(person.Age)})
	}
}
