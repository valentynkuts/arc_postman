package main

import (
	"net/http"
	"postman/go_firestore_books/books"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/books", books.Index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/books/show", books.Show)
	http.HandleFunc("/books/create", books.Create)
	http.HandleFunc("/books/create/process", books.CreateProcess)
	http.HandleFunc("/books/update", books.Update)
	http.HandleFunc("/books/update/process", books.UpdateProcess)
	http.HandleFunc("/books/delete/process", books.DeleteProcess)

	//http.HandleFunc("/books/add/json", books.MybookFromJson)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
