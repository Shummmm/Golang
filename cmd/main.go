package main

import (
	l "Library/library"
	"fmt"
)

func main() {
	lib := l.Library{
		Books: []l.Book{{ //здесь нужны двойные скобки - потому что идет слайс
			ID:        1,
			Author:    "Гоголь",
			Year:      1872,
			Title:     "Вечера на хуторе близ Диканьки",
			Available: true,
		},
			{
				ID:        2,
				Author:    "Пушкин",
				Year:      1810,
				Title:     "Онегин",
				Available: true,
			},

			{ID: 3,
				Author:    "Пушкин",
				Year:      1810,
				Title:     "Песня о царе Салтане",
				Available: true,
			},
			{
				ID:        4,
				Author:    "Тургенев",
				Year:      1901,
				Title:     "Записки охотника",
				Available: true,
			},
		},
		Users: []l.User{{
			ID:             1,
			Name:           "Русанов Леонид",
			BorrowedBookId: []int{}},
			{
				ID:             2,
				Name:           "Гусак Евгений",
				BorrowedBookId: []int{}},
			{ID: 3,
				Name:           "Акатов Анатолий",
				BorrowedBookId: []int{}},
		},
	}
	fmt.Println(lib.GetAvailableBooks())
	//fmt.Println(lib.Users[1].Borrow(1))
	//fmt.Println(l.BorrowBook(&lib, lib.Users[].ID), lib.Books[1].ID) // плохо работает
	//fmt.Println(lib.GetAvailableBooks())  // книга ушла
	//fmt.Println(lib.GetBooksByAuthor("Пушкин"))//работает
	//fmt.Println(lib.Users)
	//l.ReturnBook(&lib, lib.Users[1].ID, lib.Books[1].ID)// я не знаю как  мне вывести именно ID, сейчас выводит айди первого индекса
	fmt.Println(l.BorrowBook(&lib, 2, 2))
	fmt.Println(lib.GetAvailableBooks()) // книга ушла
	//fmt.Println(l.ReturnBook(&lib, 2, 2))
	fmt.Println(lib.GetAvailableBooks()) // книга вернулась

	/*Identifiable := []l.Library{
		{Books: []l.Book{{
			ID:        5,
			Author:    "Tolkin",
			Year:      1910,
			Title:     "The Lord of the Rings",
			Available: true,
		},
			{ID: 25,
				Author:    "Tolkin",
				Year:      1915,
				Title:     "The Lord of the Rings: return of the king",
				Available: true},},
		Users: []l.User{{
			Name:           "Бычков Сергей",
			ID:             10,
			BorrowedBookId: []int{},
		}},
	}}*/

	//Identifiable := []l.Library{
	//	{Books: []l.Book{{ID: 5, Author:"Tolkin",Year: 1910, Title: "The Lord of the Rings", Available: true},},}, {Users: []l.User{{ID: 25, Name: "Бычков Сергей", BorrowedBookId: []int{}}}}},

	var identifiableItems []l.Identifiable
	// Добавляем книги в срез интерфейсов
	for _, book := range lib.Books {
		identifiableItems = append(identifiableItems, book)
	}
	// Добавляем пользователей в срез интерфейсов
	for _, user := range lib.Users {
		identifiableItems = append(identifiableItems, user)
	}

	for _, item := range identifiableItems {
		fmt.Println(item.GetID())
	}
	l.PrintBooks(lib.Books)

	fmt.Println(lib.FindBookByID(1))
	fmt.Println(lib.FindUserByID(2))

}
