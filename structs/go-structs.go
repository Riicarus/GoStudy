package structs

import (
	"fmt"
)

// declare one type using key word "type"
// type myint int

// declare one complex type using "type" and "struct"
// struct is value pass, we may use pointer
type Book struct {
	title string
	auth  string
}

func (book *Book) GetAuth() string {
	// this is a value copy of the invoking object
	return book.auth
}

func (book *Book) SetAuth(newAuth string) {
	book.auth = newAuth
}

type StoryBook struct {
	// this means StoryBook is a sub struct of Book
	Book
	SuitableAge int
}

// override super method
func (book *StoryBook) SetAuth(newAuth string) {
	book.auth = "StoryBookAuth: " + newAuth
}

// add new method
func (book *StoryBook) GetSuitableAge() int {
	return book.SuitableAge
}

func LearnStructs() {
	fmt.Println("---------- structs ----------")

	var book1 Book
	book1.title = "Golang"
	book1.auth = "riicarus"

	fmt.Printf("%v\n", book1)

	changeBook(&book1)
	fmt.Printf("%v\n", book1)

	book2 := Book{auth: "riicarus", title: "Java"}
	// book2 := Book{"riicarus", "Java"} is also usable
	fmt.Println(book2.GetAuth())
	book2.SetAuth("skyline")
	fmt.Println(book2.GetAuth())

	// some interesting usage
	fmt.Println((*Book).GetAuth(&book2))

	// create sub struct
	storyBook1 := StoryBook{Book{"riicarus", "cpp"}, 18}
	storyBook1.SetAuth("skyline")
	fmt.Println(storyBook1.GetSuitableAge())
	fmt.Println(storyBook1.GetAuth())

	fmt.Println("")
}

func changeBook(book *Book) {
	book.auth = "skyline"
}