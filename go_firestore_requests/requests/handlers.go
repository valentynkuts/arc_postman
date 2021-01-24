package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func DoUserReq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//dump from real request to custom MyRequest
	req, err := ReqToMyReq(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}
	fmt.Println("from ReqToMyReq: ",req)

	// add to firestore
	/*err =  AddMyReq(req)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}
	fmt.Println("from AddMyReq: ",req)*/

	//todo Goroutines
	go GoAddMyReq(req)

	fmt.Println(req)
	fmt.Println(req.UserId)
	fmt.Println(req.Method)
	fmt.Println(req.Host)
	fmt.Println(req.Url)
	fmt.Println(req.Headers) //nil
	fmt.Println(req.Params)  //nil
	fmt.Println(req.Body)    //nil  todo check for nill

	// make Client to do request
	url := "http://" + req.Host + req.Url
	fmt.Println("Url to do request: ", url)

	// ok
	if req.Method == "POST" {

		if req.Body == nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}


		//body - bytes of json
		client("POST", url, req.Headers, req.Body, req.Params)
	}
	// ok
	if req.Method == "GET" {

		body := client("GET", url, req.Headers, req.Body, req.Params)

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
	// ok
	if req.Method == "PUT" {

		if req.Body == nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		b := client("PUT", url, req.Headers, req.Body, req.Params)

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
	// ok
	if req.Method == "DELETE" {
		client("DELETE", url, req.Headers, req.Body, req.Params)
	}
}

func GetUserReqs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	userId := ps.ByName("userId") //id of request
	fmt.Println("userId - ", userId)

	//get host,url,method, headers, params, body from firestore by id
	reqs, err := GetUserRequests(userId)
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(reqs)

	jsonReqs, _ := json.Marshal(reqs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonReqs)
}

func PostReq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//add requests to firestore
	req, err := AddRequest(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}
	fmt.Println(req)

}
//add requests to firestore with id :reqId
func PostReqWithId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reqID := ps.ByName("reqId")
	fmt.Println("param - ", reqID)

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}


	//add requests to firestore with id
	req, err := AddRequestWithId(r, reqID)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}
	fmt.Println(req)

}

//get request from firestore by id
//make Client to do request from firestore
//
func GetReqWithId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	reqId := ps.ByName("reqId") //id of request
	fmt.Println("param - ", reqId)

	//get host,url,method, headers, params, body from firestore by id
	req, err := GetRequestWithId(reqId)
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(req)
	fmt.Println(req.UserId)
	fmt.Println(req.Method)
	fmt.Println(req.Host)
	fmt.Println(req.Url)
	fmt.Println(req.Headers) //nil
	fmt.Println(req.Params)  //nil
	fmt.Println(req.Body)    //nil  todo check for nill

	url := "http://" + req.Host + req.Url
	fmt.Println(url)

    // ok
	if req.Method == "POST" {

		//body, err := ioutil.ReadAll(r.Body)

		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusBadRequest)
		//}

		if req.Body == nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}


		//body - bytes of json
		client("POST", url, req.Headers, req.Body, req.Params)
	}
     // ok
	if req.Method == "GET" {

		body := client("GET", url, req.Headers, req.Body, req.Params)

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
   // ok
	if req.Method == "PUT" {

		//body, err := ioutil.ReadAll(r.Body)

		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusBadRequest)
		//}

		if req.Body == nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		b := client("PUT", url, req.Headers, req.Body, req.Params)

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
    // ok
	if req.Method == "DELETE" {
		client("DELETE", url, req.Headers, req.Body, req.Params)
	}

	//----------
}
// todo return error
func client(
	method string,
	url string,
	myHeaders map[string]interface{},
	myBody map[string]interface{},
	myParams map[string]interface{}) []byte {

	fmt.Println("URL:>", url)

	//---- body ----
	var b []byte
	var err error
	if myBody != nil{
		b, err = json.Marshal(myBody)
		if err != nil {
			panic(err)
		}
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	//---- headers ----
   if myHeaders != nil{
	   for key, value := range myHeaders {
		   req.Header.Set(key, value.(string))
	   }
   }

	//---- params ----
	if myParams != nil {
		q := req.URL.Query()
		for key, value := range myParams {
			q.Add(key, value.(string))
		}
		req.URL.RawQuery = q.Encode()
	}

	fmt.Println(req.URL.String())

	//--------------------
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return body
}
// get all id of requests
func GetAllIdReq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	IDs, err := AllIdRequest()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonIDs, _ := json.Marshal(IDs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonIDs)


}
//-----------------
func GetReq1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	paramReq := ps.ByName("req")
	fmt.Println("param - ", paramReq)

	paramId := ps.ByName("id")
	fmt.Println("param - ", paramId)

	//get host,url,method from firestore
	req, err := OneRequest(paramReq)
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	//fmt.Println(req)
	//fmt.Println(req.Method)
	//fmt.Println(req.Host)
	//fmt.Println(req.Url)

	bookID := "/" + paramId
	url := "http://" + req.Host + req.Url + bookID
	fmt.Println(url)

	if req.Method == "POST" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		//body - bytes of json
		client1("POST", url, body)
	}

	if req.Method == "GET" {

		body := client1("GET", url, nil)

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}

	if req.Method == "PUT" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		b := client1("PUT", url, body)

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}

	if req.Method == "DELETE" {
		client1("DELETE", url, nil)
	}

}

func client1(method string, url string, json []byte) []byte {
	fmt.Println("URL:>", url)
	fmt.Println(bytes.NewBuffer(json))

	req, err := http.NewRequest(method, url, bytes.NewBuffer(json))
	req.Header.Set("X-Custom-Header", "myclient")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return body
}
