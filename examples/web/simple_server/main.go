package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path:", r.URL.Path)
	fmt.Fprintf(w, "hello go")
}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
