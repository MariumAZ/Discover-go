package main

import "fmt"

type Book struct { 

	Title string

}

func GetTitle(b Book) string {

	return b.Title
}

func SetTitle( b *Book, title string) {

	(*b).Title= title 
}

func main() {

	b := Book{Title: "For the Love of Go "}
	fmt.Println(GetTitle(b))
	SetTitle( &b, "Omnia")
	fmt.Println(GetTitle(b))


}