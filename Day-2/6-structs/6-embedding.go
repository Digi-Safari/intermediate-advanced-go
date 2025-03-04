package main

import "fmt"

type User struct {
	name  string
	email string
}

func (u *User) UpdateEmail(email string) {
	u.email = email
}

func (u *User) Print() {
	fmt.Println(u)
}

type BookAuthor struct {
	// anonymous field name is the same as type name
	User  // embedding // anonymous field// field with no name // in embedding we just write the type name
	bio   string
	books []string
}

func (b *BookAuthor) Print() {
	fmt.Println(b)
}

func (b *BookAuthor) AddBook(book string) {
	b.books = append(b.books, book)
}

func main() {
	u := User{
		name:  "Rahul",
		email: "rahul@gmail.com",
	}

	b := BookAuthor{
		User: u,
		bio:  "random bio",
	}

	b.AddBook("book1")
	b.Print() // print method of the author type
	b.UpdateEmail("rahul@yahoo.com")
	b.User.Print() // print method of the user type
}
