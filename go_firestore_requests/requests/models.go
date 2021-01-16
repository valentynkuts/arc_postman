package requests

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"postman/go_firestore_requests/config"
)

type Req struct {
	//ID     string
	Method string
	Host   string
	Url    string
}

type MyRequest struct {
	//ID     string
	Method  string                 `firestore:"method,omitempty"`
	Host    string                 `firestore:"host,omitempty"`
	Url     string                 `firestore:"url,omitempty"`
	Headers map[string]interface{} `firestore:"headers,omitempty"`
	Params  map[string]interface{} `firestore:"params,omitempty"`
	Body    map[string]interface{} `firestore:"body,omitempty"` // json
	//Headers  map[string]string `firestore:"headers,omitempty"`
	//Params map[string]string `firestore:"params,omitempty"`
	//Body map[string]interface{} `firestore:"body,omitempty"`  // json
}

func OneRequest(str string) (Req, error) {
	req := Req{}
	ID := str
	fmt.Println("str - ", ID)
	if ID == "" {
		return req, errors.New("400. Bad Request.")
	}

	ctx := context.Background()
	doc, _ := config.Client.Collection("requests").Doc(ID).Get(ctx)

	fmt.Println("OneRequest --", doc.Data())

	req = Req{
		Method: doc.Data()["method"].(string),
		Host:   doc.Data()["host"].(string),
		Url:    doc.Data()["url"].(string),
	}

	return req, nil
}

func GetRequest(str string) (MyRequest, error) {
	req := MyRequest{}
	ID := str
	fmt.Println("str - ", ID)
	if ID == "" {
		return req, errors.New("400. Bad Request.")
	}

	ctx := context.Background()
	doc, _ := config.Client.Collection("requests").Doc(ID).Get(ctx)

	fmt.Println("GetRequest --", doc.Data())

	req = MyRequest{
		Method: doc.Data()["method"].(string),
		Host:   doc.Data()["host"].(string),
		Url:    doc.Data()["url"].(string),
		//Headers :  doc.Data()["headers"].(map[string]interface{}),
		//Params: doc.Data()["params"].(map[string]interface{}),
		//Body :doc.Data() ["body"].(map[string]interface{}),
	}

	if val, ok := doc.Data()["headers"]; ok {
		//req.Headers = doc.Data()["headers"].(map[string]interface{})
		req.Headers = val.(map[string]interface{})
	}

	if val, ok := doc.Data()["params"]; ok {
		//req.Headers = doc.Data()["headers"].(map[string]interface{})
		req.Params = val.(map[string]interface{})
	}

	if val, ok := doc.Data()["body"]; ok {
		//req.Headers = doc.Data()["headers"].(map[string]interface{})
		req.Body = val.(map[string]interface{})
	}

	return req, nil
}

// json to map[string]interface{}
func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}

//Add request to firestore
func AddRequest(r *http.Request, reqID string) (MyRequest, error) {
	//req1 := Req1{}
	//ID:= str
	//fmt.Println("str - ",ID)
	jsonMap := make(map[string]interface{})
	//err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	err := json.NewDecoder(r.Body).Decode(&jsonMap)
	if err != nil {
		panic(err)
	}
	dumpMap("", jsonMap)

	fmt.Println("jsonMap - ", jsonMap)
	fmt.Println("------------------------------------ ")
	//var res map[string]interface{}

	//res = jsonMap

	//fmt.Println("res -  ", res)

	fmt.Println("Method -  ", jsonMap["Method"])
	fmt.Println("Host -  ", jsonMap["Host"])
	fmt.Println("Url -  ", jsonMap["Url"])
	fmt.Println("Headers -  ", jsonMap["Headers"])
	fmt.Println("Params -  ", jsonMap["Params"])
	fmt.Println("Body -  ", jsonMap["Body"])

	//if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || bk.Price == "" {
	//	return req1, errors.New("400. Bad Request.")
	//}

	req := MyRequest{
		Method:  jsonMap["Method"].(string),
		Host:    jsonMap["Host"].(string),
		Url:     jsonMap["Url"].(string),
		Headers: jsonMap["Headers"].(map[string]interface{}),
		Params:  jsonMap["Params"].(map[string]interface{}),
		Body:    jsonMap["Body"].(map[string]interface{}),
	}

	if req.Method == "" || req.Host == "" || req.Url == "" || req.Headers == nil {
		return req, errors.New("400. Bad Request.")
	}

	ctx := context.Background()
	_, err = config.Client.Collection("requests").Doc(reqID).Set(ctx, req)

	//_, _, err = config.Client.Collection("books").Add(context.Background(),
	//	map[string]interface{}{
	//		"Isbn":   bk.Isbn,
	//		"Title":  bk.Title,
	//		"Author": bk.Author,
	//		"Price":  bk.Price,
	//	})
	//
	//if err != nil {
	//	log.Fatalf("Failed to add a new book: %w", err)
	//	//fmt.Errorf("Failed to iterate the list of requests: %w", err)
	//
	//}
	//   "Content-Type": "application/json",

	return req, nil
}

// get all id of requests from firestore
func AllRequest() ([]string, error) {
	var IDs []string
	q := config.Client.Collection("requests")
	docs, err := q.Documents(context.Background()).GetAll()
	if err != nil {
		fmt.Print(err)
	}
	for _, doc := range docs {
		fmt.Println(doc.Data())

    	ID := doc.Ref.ID
		IDs = append(IDs, ID)
	}

	return IDs, nil

}