package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*type Articles struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	id      string `json:"id"`
	message string `json:"message"`
}*/

func handler(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://localhost:9000/")
	//data := json.RawMessage(string(resp)) // declaring a data string parsing data as raw

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).decode(resp)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//var generic map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	text := string(body)
	json.Unmarshal([]byte(body), &resp)
	str := fmt.Sprint(text) //converting data type interface to string
	str1 := Reverse(str)
	fmt.Printf(string(str1))

}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are on the home page")
}

func main() {

	r := mux.NewRouter() //.StrictSlash(true)
	r.HandleFunc("/articles", handler).Methods("GET")
	r.HandleFunc("/", homepage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

}
