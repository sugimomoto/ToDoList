package main

import (
	"ToDoList/data"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello")
	http.HandleFunc("/", indexTemplate)
	http.ListenAndServe(":8081", nil)

	fmt.Println("End")
}

func indexTemplate(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}

	ToDoList, err := data.ReadToDoList()

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(writer, ToDoList)
}
