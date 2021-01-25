package requests

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"net/http"
	"postman/go_firestore_requests/config"
)

type MyRequest struct {
	UserId  string                 `firestore:"userId,omitempty"`
	Method  string                 `firestore:"method,omitempty"`
	Host    string                 `firestore:"host,omitempty"`
	Url     string                 `firestore:"url,omitempty"`
	Headers map[string]interface{} `firestore:"headers,omitempty"`
	Params  map[string]interface{} `firestore:"params,omitempty"`
	Body    map[string]interface{} `firestore:"body,omitempty"` // json
}


func GetRequestWithId(str string) (MyRequest, error) {
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
		UserId: doc.Data()["userId"].(string),
		Method: doc.Data()["method"].(string),
		Host:   doc.Data()["host"].(string),
		Url:    doc.Data()["url"].(string),
	}

	if val, ok := doc.Data()["headers"]; ok {
		req.Headers = val.(map[string]interface{})
	}

	if val, ok := doc.Data()["params"]; ok {
		req.Params = val.(map[string]interface{})
	}

	if val, ok := doc.Data()["body"]; ok {
		req.Body = val.(map[string]interface{})
	}

	return req, nil
}

// json to map[string]interface{}
func dumpMap1(space string, m map[string]interface{}) {
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
//ok   json to map[string]interface{}
func dumpMap(space string, m map[string]interface{}) {
	for _, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			dumpMap(space+"\t", mv)
		}
	}
}
//Dump from real request to custom MyRequest
func ReqToMyReq(r *http.Request) (MyRequest, error) {
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&jsonMap)
	if err != nil {
		panic(err)
	}
	dumpMap("", jsonMap)

	req := MyRequest{
		UserId:  jsonMap["UserId"].(string),
		Method:  jsonMap["Method"].(string),
		Host:    jsonMap["Host"].(string),
		Url:     jsonMap["Url"].(string),
		Headers: jsonMap["Headers"].(map[string]interface{}),
		Params:  jsonMap["Params"].(map[string]interface{}),
		Body:    jsonMap["Body"].(map[string]interface{}),
	}

	if req.UserId == "" || req.Method == "" || req.Host == "" || req.Url == "" || req.Headers == nil {
		return req, errors.New("400. Bad Request.")
	}

	return req, nil

}
//Add request to firestore with id
func AddRequestWithId(r *http.Request, reqID string) (MyRequest, error) {

	//dump from real request to custom MyRequest
	req, err := ReqToMyReq(r)
	if err != nil {
		return req, err
	}

	ctx := context.Background()
	_, err = config.Client.Collection("requests").Doc(reqID).Set(ctx, req)

	return req, err
}

//Add custom request (MyRequest)to firestore with random id
func AddMyReq(mr MyRequest)  error {

	if mr.UserId == "" || mr.Method == "" || mr.Host == "" || mr.Url == "" || mr.Headers == nil {
		return errors.New("400. Bad Request.")
	}

	ctx := context.Background()

	_, _, err := config.Client.Collection("requests").Add(ctx,
		map[string]interface{}{
			"userId":   mr.UserId ,
			"method":   mr.Method ,
			"host":   mr.Host ,
			"url":   mr.Url ,
			"headers":   mr.Headers ,
			"params":   mr.Params ,
			"body":   mr.Body ,

		})

	if err != nil {
		log.Fatalf("Failed to add a new Request: %w", err)
	}

	return  err

}

//Add custom request (MyRequest)to firestore with random id using Goroutines
func GoAddMyReq(mr MyRequest) {

	if mr.UserId == "" || mr.Method == "" || mr.Host == "" || mr.Url == "" || mr.Headers == nil {
		//return errors.New("400. Bad Request.")
		log.Fatalf("400. Bad Request")
	}

	ctx := context.Background()

	_, _, err := config.Client.Collection("requests").Add(ctx,
		map[string]interface{}{
			"userId":   mr.UserId ,
			"method":   mr.Method ,
			"host":   mr.Host ,
			"url":   mr.Url ,
			"headers":   mr.Headers ,
			"params":   mr.Params ,
			"body":   mr.Body ,

		})

	if err != nil {
		log.Fatalf("Failed to add a new Request: %w", err)
	}

}


//Add request to firestore with random id
func AddRequest(r *http.Request) (MyRequest, error) {

	req, err := ReqToMyReq(r)
	if err != nil {
		return req, err
	}

	ctx := context.Background()
	_, _, err = config.Client.Collection("requests").Add(ctx,
		map[string]interface{}{
			"userId":   req.UserId ,
			"method":   req.Method ,
			"host":   req.Host ,
			"url":   req.Url ,
			"headers":   req.Headers ,
			"params":   req.Params ,
			"body":   req.Body ,

		})

	if err != nil {
		log.Fatalf("Failed to add a new Request: %w", err)
	}

	return req, err
}


// get all id of requests from firestore
func AllIdRequest() ([]string, error) {
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

func GetUserRequests(userId string) ([]MyRequest, error) {
	var reqs []MyRequest
	fmt.Println("userId - ", userId)
	if userId == "" {
		return reqs, errors.New("400. Bad Request.")
	}

	ctx := context.Background()
	q := config.Client.Collection("requests").Where("userId", "==", userId)
	iter := q.Documents(ctx)
	//defer iter.Stop()
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

		req := MyRequest{
			UserId:  doc.Data()["userId"].(string),
			Method:  doc.Data()["method"].(string),
			Host:    doc.Data()["host"].(string),
			Url:     doc.Data()["url"].(string),
		}

		if val, ok := doc.Data()["headers"]; ok {
			req.Headers = val.(map[string]interface{})
		}

		if val, ok := doc.Data()["params"]; ok {
			req.Params = val.(map[string]interface{})
		}

		if val, ok := doc.Data()["body"]; ok {
			req.Body = val.(map[string]interface{})
		}

		reqs = append(reqs, req)

	}

	return reqs, nil
}
