package main

import "fmt"

type Author struct {
	Name  string
	Books []string
}

// AddBook appends a new book to the Books slice
func (a *Author) AddBook(book string) {
	a.Books = append(a.Books, book)
}

func (a *Author) PrintDetails() {
	fmt.Printf("Author Name: %s\nBooks: %v\n", a.Name, a.Books)
}

// AppendBooksToAuthor accepts the Author struct and appends values to the Books slice
func AppendBooksToAuthor(author *Author, booksToAdd []string) {
	author.Books = append(author.Books, booksToAdd...)
}

// AppendToBooks accepts the Books slice and appends some more books
func AppendToBooks(books *[]string, booksToAdd []string) {

	// ... unpacking a slice
	*books = append(*books, booksToAdd...) // Dereferencing the pointer
}

func main() {

	author := Author{Name: "John Doe", Books: []string{"Book 1", "Book 2"}}

	// Add a book using the method
	author.AddBook("Book 3")
	author.PrintDetails()

	// Append more books using a function that accepts the struct
	AppendBooksToAuthor(&author, []string{"Book 4", "Book 5"})
	author.PrintDetails()

	// Append more books using a function that accepts the Books slice
	AppendToBooks(&author.Books, []string{"Book 6", "Book 7"})
	author.PrintDetails()
}
