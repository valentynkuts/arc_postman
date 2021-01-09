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
	router.GET("/requests/:req/:id", requests.GetReq)
	//router.GET("/requests/:req", requests.GetReq)

	log.Fatal(http.ListenAndServe(":8082", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}


