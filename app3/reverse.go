package main
package stringutil

import (
	"encoding/json" // import encoding/json package
	"fmt"           // import formatted I/O
	"log"           // import logging
	"net/http"      // import HTTP implem.
)

func main() {
	resp, err := http.Get("http://localhost:9000/articles")
	if err != nil {
		log.Fatal(err)
	}
	var generic map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&generic)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generic)
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(runes)
}
