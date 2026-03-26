package library

import "fmt"

type Book struct {
	ID, Year      int
	Title, Author string
	Available     bool
}

func (b Book) GetInfo() string {
	return fmt.Sprintf("ID: %d, Title: %s, Author: %s Year: %d Available: %t", b.ID, b.Title, b.Author, b.Year, b.Available)
}

func (b Book) GetID() int { //для интерфейса
	return b.ID
}
