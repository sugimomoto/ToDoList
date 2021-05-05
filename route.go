package main

import (
	"ToDoList/data"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

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

func createFormTemplate(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("templates/create_form.html")
	if err != nil {
		panic(err)
	}
	t.Execute(writer, nil)
}

func createToDoTemplate(writer http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	subject := request.PostFormValue("subject")
	priority := request.PostFormValue("priority")

	todo := data.ToDo{
		Subject:  subject,
		Priority: priority,
	}

	_, err = todo.CreateToDo()

	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(writer, request, "/", 302)
}

func editFormTemplate(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	id := vals.Get("id")

	t, err := template.ParseFiles("templates/edit_form.html")
	if err != nil {
		panic(err)
	}

	i, _ := strconv.Atoi(id)

	ToDo, err := data.ReadToDo(i)

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(writer, ToDo)
}

func editToDoTemplate(writer http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	id := request.PostFormValue("todoid")
	subject := request.PostFormValue("subject")
	priority := request.PostFormValue("priority")

	i, _ := strconv.Atoi(id)

	todo := data.ToDo{
		Id:       i,
		Subject:  subject,
		Priority: priority,
	}

	err = todo.UpdateToDo()

	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(writer, request, "/", 302)
}
