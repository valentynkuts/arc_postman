package requests

import (
	"context"
	"errors"
	"fmt"
	"postman/go_firestore_requests/config"
)

type Req struct {
	//ID     string
	Method   string
	Host  string
	Url string
}

func OneRequest(str string) (Req, error) {
	req := Req{}
	ID:= str
	fmt.Println("str - ",ID)
	if ID == "" {
		return req, errors.New("400. Bad Request.")
	}

	ctx := context.Background()
	doc, _ := config.Client.Collection("requests").Doc(ID).Get(ctx)

	fmt.Println("OneRequest --",doc.Data())

	req = Req{
		Method:   doc.Data()["method"].(string),
		Host:  doc.Data()["host"].(string),
		Url: doc.Data()["url"].(string),
	}

	return req, nil
}
