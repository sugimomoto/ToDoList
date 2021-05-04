package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
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

	ToDoList := GetToDoLost()

	t.Execute(writer, ToDoList)
}

func GetToDoLost() []ToDo {
	var toDoList []ToDo

	for i := 0; i < 10; i++ {
		todo := ToDo{
			Id:        i,
			Subject:   "Sample",
			Priority:  "A",
			CreatedAt: time.Now(),
		}
		toDoList = append(toDoList, todo)
	}

	return toDoList
}

type ToDo struct {
	Id        int
	Subject   string
	Priority  string
	CreatedAt time.Time
}
