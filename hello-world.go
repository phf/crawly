package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!\n")
}

func main() {
	fmt.Println("Hello World!")
	http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(":3333", nil))
}
