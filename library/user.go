package library

import (
	"fmt"
)

type User struct {
	ID             int
	Name           string
	BorrowedBookId []int //ID книг которые взял
}

func (u *User) Borrow(bookID int) {
	u.BorrowedBookId = append(u.BorrowedBookId, bookID)
	//fmt.Println(u.BorrowedBookId)
}
func (u *User) Return(bookID int) {
	// найти элемент, определить его идент, и потом через индекс удалить
	for i := 0; i < len(u.BorrowedBookId); i++ {
		if u.BorrowedBookId[i] == bookID {
			u.BorrowedBookId = append(u.BorrowedBookId[:i], u.BorrowedBookId[i+1:]...) // если находится нужная книга, запоминается ее элемент и удаляется
			return
		} else {
			fmt.Println("Книга не найдена")
		}
	}
}
func (u User) HasBook(bookID int) bool {
	for i := 0; i < len(u.BorrowedBookId); i++ {
		if u.BorrowedBookId[i] == bookID {
			return true
		}
	}
	return false
}

func (u User) GetID() int { //для интерфейса
	return u.ID
}
