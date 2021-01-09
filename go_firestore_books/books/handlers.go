package books

import (
	"encoding/json"
	"net/http"
	"postman/go_firestore_books/config"
)

func Index(w http.ResponseWriter, r *http.Request) {
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
func Create(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "create.gohtml", nil)
}

func CreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := PutBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	config.TPL.ExecuteTemplate(w, "created.gohtml", bk)
}
//for test json
func MybookFromJson(w http.ResponseWriter, r *http.Request){

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

	_, err = AddBookJson(bk)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
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

func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := OneBook(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "show.gohtml", bk)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := OneBook(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "update.gohtml", bk)
}

func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := UpdateBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
		return
	}

	config.TPL.ExecuteTemplate(w, "updated.gohtml", bk)
}

func DeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := DeleteBook(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
