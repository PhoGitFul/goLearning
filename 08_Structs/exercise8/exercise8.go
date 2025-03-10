// Create a Book struct with the following fields:
// title (string)
// author (string)
// pages (int)
// published (bool)
// Create two instances of the Book struct:
// book1 with the title "To Kill a Mockingbird", author "Harper Lee", 281 pages, and published true
// book2 with the title "1984", author "George Orwell", 328 pages, and published true
// Write a function printBook that takes a Book struct as an argument and prints out the book's title, author, and number of pages.

package main

import "fmt"

type Book struct {
	title     string
	author    string
	pages     int
	published bool
}

func main() {
	var book1 Book
	var book2 Book
	// Book1 specification
	book1.title = "To Kill a Mockingbird"
	book1.author = "Harper Lee"
	book1.pages = 281
	book1.published = true
	// Book2 specification
	book2.title = "1984"
	book2.author = "George Orwell"
	book2.pages = 328
	book2.published = true
	// Print Book1 details by calling a function
	printBook(book1)
	// Print Book2 details by calling a function
	printBook(book2)
}
func printBook(bk Book) {
	fmt.Println("Title: ", bk.title)
	fmt.Println("Author: ", bk.author)
	fmt.Println("Pages: ", bk.pages)
	fmt.Println()
}
