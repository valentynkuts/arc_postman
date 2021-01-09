package requests

import (
	"bytes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func GetReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
		client("POST",url, body)
	}

	if req.Method == "GET" {


		body := client("GET",url, nil)

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}

	if req.Method == "PUT" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		b := client("PUT",url, body)

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}

	if req.Method == "DELETE" {
		client("DELETE",url, nil)
	}


}

func client(method string, url string, json []byte) []byte{
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
