package requests
/*
import (
	"bytes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)



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

	fmt.Println(req)
	fmt.Println(req.Method)
	fmt.Println(req.Host)
	fmt.Println(req.Url)

    if req.Method == "POST" {
		bookID := "/" + paramId
		url := "http://" + req.Host + req.Url + bookID
		fmt.Println(url)

		//---- client ----

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		//json := `{"Isbn":"2bk.Isbn","Title":"2bk.Title","Author":"2bk.Author","Price":"2bk.Price"}`
		//clientPost(url, json)
		clientPostByte(url, body)
	}

	if req.Method == "GET" {
		bookID := "/" + paramId
		url := "http://" + req.Host + req.Url + bookID
		fmt.Println(url)

		body := clientGet(url)

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}

	if req.Method == "PUT" {
		bookID := "/" + paramId
		url := "http://" + req.Host + req.Url + bookID
		fmt.Println(url)

		//---- client ----

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		b := client("PUT",url, body)

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}

	if req.Method == "DELETE" {
		bookID := "/" + paramId
		url := "http://" + req.Host + req.Url + bookID
		fmt.Println(url)

		client("DELETE",url, nil)

	}


}

func client(method string, url string, json []byte) []byte{
	fmt.Println("URL:>", url)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(json))
	req.Header.Set("X-Custom-Header", "myvalue")
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

func clientPostByte(url string, json []byte) {
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	req.Header.Set("X-Custom-Header", "myvalue")
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


}

func clientPost(url string, json string) {
	fmt.Println("URL:>", url)

	jsonStr := []byte(json)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
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
}

func clientGet(url string) []byte{
	fmt.Println("URL:>", url)


	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Custom-Header", "myBook")

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
}*/