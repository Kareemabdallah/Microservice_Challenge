package main

import (
	"encoding/json" // import encoding/json package
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux" // HTTP router library
)

func GetArticles(w http.ResponseWriter, r *http.Request) {

	// Open our jsonFile as a byte array
	byteValue, _ := ioutil.ReadFile("db.json")

	// declaring a data string parsing data as raw
	data := json.RawMessage(string(byteValue))

	// defining map function with keys as strings and values as type.
	var articles map[string]*json.RawMessage

	// unmarshalling byteArray which contains jsonFile's content into 'articles'
	err := json.Unmarshal(data, &articles)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// print articles content
	fmt.Println(articles)

	// encoding our articles array into a JSON string
	json.NewEncoder(w).Encode(articles)
}

func main() {

	// Defining our HTTP request multiplexer
	router := mux.NewRouter()

	// Registing URL path /articles where acquiring all articles
	router.HandleFunc("/articles", GetArticles).Methods("GET")

	// printing Listening on localhost:9000
	log.Println("Listening on localhost:9000")

	// listening on port 9000
	log.Fatal(http.ListenAndServe(":9000", router))
}
