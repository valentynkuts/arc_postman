package books

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"postman/go_firestore_books_julienschimdt/config"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bks, err := AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "books.gohtml", bks)
}

func AddMybookFromJsonWithID(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	bookID := ps.ByName("bookID")
	fmt.Println("param - ", bookID)

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	//---- request ----
	bk := Book{}
	err := json.NewDecoder(r.Body).Decode(&bk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = AddJsonBookWithID(bookID, bk)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	//---- response ----//todo delete till end of func ↓
	js, err := json.Marshal(bk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetMybookWithID(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	bookID := ps.ByName("bookID")
	fmt.Println("param - ", bookID)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := GetBookWithID(bookID)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	//---- response ----
	js, err := json.Marshal(bk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func UpdateMybookFromJsonWithID(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	bookID := ps.ByName("bookID")
	fmt.Println("param - ", bookID)

	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	//---- request ----
	bk := Book{}
	err := json.NewDecoder(r.Body).Decode(&bk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = UpdateJsonBookWithID(bookID, bk)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	//---- response ----//todo delete till end of func ↓
	js, err := json.Marshal(bk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func DeleteMybookWithID(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	bookID := ps.ByName("bookID")
	fmt.Println("param - ", bookID)

	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := DeleteBookWithID(bookID)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

}
