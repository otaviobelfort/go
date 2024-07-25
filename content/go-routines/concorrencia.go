package main

import (
	"fmt"
	"net/http"
)

var number uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		w.Write([]byte(fmt.Sprintf("Você é o visitante de número: %d ", number)))
	})
	http.ListenAndServe(":8080", nil)

}
