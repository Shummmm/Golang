package library

import "fmt"

func BorrowBook(lib *Library, userID, bookID int) string {
	//Ищу пользователя по ID
	var targetUser *User
	for i := 0; i < len(lib.Users); i++ {
		if lib.Users[i].ID == userID {
			targetUser = &lib.Users[i]
			break
		}
	}
	//Проверка, если нет пользователя
	if targetUser == nil {
		return "Ошибка: пользователь не найден"
	}
	// Ищу книгу по ID
	var targetBook *Book

	for i := 0; i < len(lib.Books); i++ {
		if lib.Books[i].ID == bookID {
			targetBook = &lib.Books[i]
			break
		}
	}
	//Проверка, если нет книги
	if targetBook == nil {
		return "Ошибка: книга не найдена"
	}
	// Если уже выдана книга
	if !targetBook.Available {
		return "Ошибка: книга уже выдана"
	}
	//Выдаю книгу
	targetUser.Borrow(bookID)
	// Меняю доступность
	targetBook.Available = false

	return "Успешно: книга выдана"
}

func ReturnBook(lib *Library, userID, bookID int) string {
	var targetUser *User
	for i := 0; i < len(lib.Users); i++ {
		if lib.Users[i].ID == userID {
			targetUser = &lib.Users[i]
			break
		}
	}

	if targetUser == nil {
		return "Oшибка: пользователь не найден"
	}

	var targetBook *Book
	for i := 0; i < len(lib.Books); i++ {
		if lib.Books[i].ID == bookID {
			targetBook = &lib.Books[i]
			break
		}
	}
	if targetBook == nil {
		return "Ошибка: книга не найдена"
	}

	/*if !targetBook.Available {
		return "Ошибка: книга уже выдана"
	}*/

	targetUser.Return(bookID)
	targetBook.Available = true
	return "Успешно: книга возвращена"

}
func PrintBooks(books []Book) { // Выводит список книг
	for i := 0; i < len(books); i++ {
		fmt.Println(books[i].GetInfo())
	}
}
