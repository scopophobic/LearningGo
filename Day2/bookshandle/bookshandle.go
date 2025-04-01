package bookshandle

import (
	"fmt"
)

type Books struct {
	book_name string
	book_author string
	book_id int
	book_genre []string
}


func Hellofrombooks() { // Notice capital "B"
    fmt.Println("Hello from bookshandle!")
}

func Bookshandle(){
	fmt.Println("Hellow world!")
	var books []Books
	var booksID []int
	for{
		var enter string;
		fmt.Println("you wanna make an entry ? (yes/no) ")
		fmt.Scan(&enter)
		if enter == "no" {break}
		var (
			bName string
			aName string
			bId int
			bGenre []string
		)
		fmt.Println("Enter Book Name :")
		fmt.Scan(&bName)
		fmt.Println("Enter Author's Name :")
		fmt.Scan(&aName)
		// fmt.Println("Enter Book ID :")
		// fmt.Scan(&bId)
		for {
			fmt.Println("Enter Book ID:")
			fmt.Scan(&bId)

			// Check if ID already exists
			exists := false
			for _, id := range booksID {
				if id == bId {
					exists = true
					break
				}
			}

			if exists {
				fmt.Println("This Book ID already exists! Enter a unique one.")
			} else {
				break
			}
		}

		fmt.Println("Enter Genre (up to 5, type -1 to stop):")
		for i:= 0 ; i<=5 ; i++{
			var temp string 
			fmt.Scan(&temp)
			if temp == "-1"  {break}
			bGenre = append(bGenre, temp)
		}

		book := Books{bName, aName, bId, bGenre}
		books = append(books, book)
		booksID = append(booksID, bId)

		fmt.Println("\nAll Books:")

		for _,book := range books{
			fmt.Printf("ID: %d | Name: %s | Author: %s | Genres: %v\n",
			book.book_id, book.book_name, book.book_author, book.book_genre)
		}


	}
}

