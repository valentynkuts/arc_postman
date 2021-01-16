package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"postman/go_firestore_requests/requests"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	//:req to get from db, :id to make client request
	//router.GET("/requests/:req/:id", requests.GetReq)

    //-------------
    //reqId - id of request in firestore
	//get request from firestore by id
	//make Client to do request that we have got from firestore
    router.GET("/requests/:reqId", requests.GetReq)
	//get all id of requests
	router.GET("/requests", requests.GetAllReq)
	// add request to firestore with id
	router.POST("/requests/:reqId", requests.PostReq)

	log.Fatal(http.ListenAndServe(":8082", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}


