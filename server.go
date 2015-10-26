package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pawel.poczta@gmail.com")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	err := http.ListenAndServe(":8000", mux)
	fmt.Println(err)
}
