package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello fadia , hiii welcome back;;!")
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
