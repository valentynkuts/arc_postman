package books

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"net/http"
	"postman/go_firestore_books/config"
)

type Book struct {
	ID     string
	Isbn   string
	Title  string
	Author string
	Price  string
}

func AllBooks() ([]Book, error) {
	var bks []Book
	//bks := []Book{}
	q := config.Client.Collection("books")
	docs, err := q.Documents(context.Background()).GetAll()
	if err != nil {
		fmt.Print(err)
	}
	for _, doc := range docs {
		fmt.Println(doc.Data())

		book := Book{
			ID: doc.Ref.ID,
			Isbn:   doc.Data()["Isbn"].(string),
			Title:  doc.Data()["Title"].(string),
			Author: doc.Data()["Author"].(string),
			//Price: doc.Data()["price"].(float32),
			Price: doc.Data()["Price"].(string),
		}

		bks = append(bks, book)
	}

	return bks, nil
}

// add book
func PutBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	bk.Price = r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || bk.Price == "" {
		return bk, errors.New("400. Bad request. All fields must be complete.")
	}

	_, _, err := config.Client.Collection("books").Add(context.Background(),
		map[string]interface{}{
			"Isbn":   bk.Isbn,
			"Title":  bk.Title,
			"Author": bk.Author,
			"Price":  bk.Price,
		})

	if err != nil {
		log.Fatalf("Failed to add a new book: %w", err)
		//fmt.Errorf("Failed to iterate the list of requests: %w", err)

	}

	return bk, nil
}

func AddBookJson(bk Book) (Book, error) {

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || bk.Price == "" {
		return bk, errors.New("400. Bad request. All fields must be complete.")
	}

	_, _, err := config.Client.Collection("books").Add(context.Background(),
		map[string]interface{}{
			"Isbn":   bk.Isbn,
			"Title":  bk.Title,
			"Author": bk.Author,
			"Price":  bk.Price,
		})

	if err != nil {
		log.Fatalf("Failed to add a new book: %w", err)
		//fmt.Errorf("Failed to iterate the list of requests: %w", err)

	}

	return bk, nil
}

//for test json
//add book ,set id = redC1
func AddJsonBookWithID(id string, bk Book) (Book, error) {

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || bk.Price == "" {
		return bk, errors.New("400. Bad request. All fields must be complete.")
	}
    ID := id
	//doc := make(map[string]interface{})
	//doc["Isbn"] = bk.Isbn
	//doc["Title"] = bk.Title
	//doc["Author"] = bk.Author
	//doc["Price"] = bk.Price

    ctx := context.Background()
	 _, err := config.Client.Collection("books").Doc(ID).Set(ctx,
		map[string]interface{}{
			"Isbn":   bk.Isbn,
			"Title":  bk.Title,
			"Author": bk.Author,
			"Price":  bk.Price,
		})

	if err != nil {
		log.Fatalf("Failed to add a new book: %w", err)
		//fmt.Errorf("Failed to iterate the list of requests: %w", err)

	}

	return bk, nil
}


func OneBook(r *http.Request) (Book, error) {
	book := Book{}
	ID:= r.FormValue("id") //from request query
	fmt.Print(ID)
	if ID == "" {
		return book, errors.New("400. Bad Request.")
	}

	ctx := context.Background()
	doc, _ := config.Client.Collection("books").Doc(ID).Get(ctx)

	fmt.Println(doc.Data())

	book = Book{
		ID: ID,
		Isbn:   doc.Data()["Isbn"].(string),
		Title:  doc.Data()["Title"].(string),
		Author: doc.Data()["Author"].(string),
		Price: doc.Data()["Price"].(string),
	}

	return book, nil
}

func OneBook1(r *http.Request) (Book, error) {
	book := Book{}
	isbn := r.FormValue("isbn") //from request query
	fmt.Print(isbn)
	if isbn == "" {
		return book, errors.New("400. Bad Request.")
	}

	ctx := context.Background()
	q := config.Client.Collection("books").Where("Isbn", "==", isbn).Limit(1)
	iter := q.Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Print(err)
			//return err
		}
		fmt.Println(doc.Data())
		fmt.Println(doc.Ref.ID)

		book = Book{
			Isbn:   doc.Data()["Isbn"].(string),
			Title:  doc.Data()["Title"].(string),
			Author: doc.Data()["Author"].(string),
			//Price: doc.Data()["price"].(float32),
			Price: doc.Data()["Price"].(string),
		}
	}

	return book, nil
}
func UpdateBook(r *http.Request) (Book, error) {
	bk := Book{}
	ID:= r.FormValue("id") //from request query
	fmt.Print(ID)
	if ID == "" {
		return bk, errors.New("400. Bad Request.")
	}
	// get form values
	bk.ID = ID
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	bk.Price = r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || bk.Price == "" {
		return bk, errors.New("400. Bad request. All fields must be complete.")
	}
    //Set or Update
	ctx:=context.Background()
    str:= "books/"+ID
    fmt.Print(str)
	_, err := config.Client.Doc(str).Set(ctx, map[string]interface{}{
		"Isbn":   bk.Isbn,
		"Title":  bk.Title,
		"Author": bk.Author,
		"Price":  bk.Price,
	})


	if err != nil {
		log.Fatalf("Failed to add a new book: %w", err)
		//fmt.Errorf("Failed to iterate the list of requests: %w", err)

	}

	return bk, nil
}

func DeleteBook(r *http.Request) error {
	ID:= r.FormValue("id")
	if ID == "" {
		return errors.New("400. Bad Request.")
	}

	ctx:=context.Background()
	_, err := config.Client.Collection("books").Doc(ID).Delete(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return errors.New("500. Internal Server Error")
	}
	return nil
}
