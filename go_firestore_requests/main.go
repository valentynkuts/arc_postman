package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"postman/go_firestore_requests/requests"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	// get request from firestore by id
	// make Client to do request that we have got from firestore
    router.GET("/requests/:reqId", requests.GetReqWithId)

	// get all id of requests
	router.GET("/requests", requests.GetAllIdReq)

	// add request to firestore with id
	router.POST("/requests/:reqId", requests.PostReqWithId)

    // add request to firestore with random id
	router.POST("/requests_add", requests.PostReq)  // todo  requests_add

	// get user's requests with userId
	router.GET("/user/requests/:userId", requests.GetUserReqs) // todo

	//Processing json from the request, selecting information
	//regarding the request client (host, url, headers, parameters, method, body),
	//storing this data in a database, making the client to execute the request
	//from the received information(host, url, ...),
	//sending the results(response) of the request back.
	router.POST("/requests", requests.DoUserReq)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}

	//log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}


