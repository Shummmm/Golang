package library

type Library struct {
	Books []Book
	Users []User
}

func (l Library) AddBook(book Book) {
	l.Books = append(l.Books, book) //добавляю книгу в слайс
}
func (l Library) AddUser(user User) {
	l.Users = append(l.Users, user) //добавляю пользователя в слайс
}
func (l Library) FindBookByID(id int) *Book { //возвращает указатель на книгу
	for i := 0; i < len(l.Books); i++ {
		if l.Books[i].ID == id {
			return &l.Books[i]
		}
	}
	return nil
}

func (l Library) FindUserByID(id int) *User { //возвращает указатель на пользователя
	for i := 0; i < len(l.Users); i++ {
		if l.Users[i].ID == id {
			return &l.Users[i]
		}
	}
	return nil
}

func (l Library) GetAvailableBooks() []Book { //возвращает все доступные книги
	var AvailableBooks []Book
	for i := 0; i < len(l.Books); i++ {
		if l.Books[i].Available == true {
			AvailableBooks = append(AvailableBooks, l.Books[i])
		}
	}
	return AvailableBooks
}

func (l Library) GetBooksByAuthor(author string) []Book { //возвращает книги автора
	var AuthorBooks []Book
	for i := 0; i < len(l.Books); i++ {
		if l.Books[i].Author == author {
			AuthorBooks = append(AuthorBooks, l.Books[i])
		}
	}
	return AuthorBooks
}
