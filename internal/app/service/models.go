package service

import "github.com/google/uuid"

type Person struct {
	UUID uuid.UUID
	Name string
	Age  int
}
