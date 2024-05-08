package main

import (
	"fmt"
)

type Book struct {
	Title       string
	Author      string
	Publication int
	IsAvailable bool
}

type Library struct {
	books []*Book
}

func (l *Library) AddBook(book *Book) {
	l.books = append(l.books, book)
	fmt.Println("Book added successfully!")
}

func (l *Library) CheckoutBook(title string) *Book {
	for i, book := range l.books {
		if book.Title == title && book.IsAvailable {
			book.IsAvailable = false
			fmt.Println("Book checked out successfully!")
			return l.books[i]
		}
	}
	fmt.Println("Book not found or unavailable.")
	return nil
}

func (l *Library) ReturnBook(title string) {
	for _, book := range l.books {
		if book.Title == title {
			book.IsAvailable = true
			fmt.Println("Book returned successfully!")
			return
		}
	}
	fmt.Println("Book not found in the library.")
}

func main() {
	myLibrary := Library{books: []*Book{}}

	var choice string

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add a book")
		fmt.Println("2. Checkout a book")
		fmt.Println("3. Return a book")
		fmt.Println("4. Exit")
		fmt.Println("Enter your choice:")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			var title, author string
			var publication int
			fmt.Println("Enter book title:")
			fmt.Scanln(&title)
			fmt.Println("Enter book author:")
			fmt.Scanln(&author)
			fmt.Println("Enter publication year:")
			fmt.Scanln(&publication)
			myLibrary.AddBook(&Book{Title: title, Author: author, Publication: publication, IsAvailable: true})
		case "2":
			var title string
			fmt.Println("Enter title of the book to checkout:")
			fmt.Scanln(&title)
			if checkedOutBook := myLibrary.CheckoutBook(title); checkedOutBook != nil {
				fmt.Printf("Book '%s' by %s (Published: %d) has been checked out.\n", checkedOutBook.Title, checkedOutBook.Author, checkedOutBook.Publication)
			}
		case "3":
			var title string
			fmt.Println("Enter title of the book to return:")
			fmt.Scanln(&title)
			myLibrary.ReturnBook(title)
		case "4":
			fmt.Println("Exiting the library system.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
