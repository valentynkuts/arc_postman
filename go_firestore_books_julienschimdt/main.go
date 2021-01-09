package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"postman/go_firestore_books_julienschimdt/books"
)

func main() {
	router := httprouter.New()

	router.GET("/", index)
	router.GET("/books", books.Index) //GET

	router.POST("/books/:bookID", books.AddMybookFromJsonWithID)
	router.GET("/books/:bookID", books.GetMybookWithID)
	router.PUT("/books/:bookID", books.UpdateMybookFromJsonWithID)
	router.DELETE("/books/:bookID", books.DeleteMybookWithID)

	log.Fatal(http.ListenAndServe(":8081", router))

}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

