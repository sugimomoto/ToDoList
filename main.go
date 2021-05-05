package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello")
	http.HandleFunc("/", indexTemplate)
	http.HandleFunc("/createform", createFormTemplate)
	http.HandleFunc("/create", createToDoTemplate)
	http.HandleFunc("/editform", editFormTemplate)
	http.HandleFunc("/edit", editToDoTemplate)
	http.ListenAndServe(":8081", nil)

	fmt.Println("End")
}
