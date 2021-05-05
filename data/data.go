package data

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:Password!@/sample?parseTime=true")

	if err != nil {
		panic(err)
	}
}

type ToDo struct {
	Id        int
	Subject   string
	Priority  string
	CreatedAt ViewerTime
}

type ViewerTime time.Time

func (v ViewerTime) String() string {
	var t time.Time = time.Time(v)

	return t.Format("2006-01-02 15:04:05")
}

func (todo *ToDo) CreateToDo() (result ToDo, err error) {
	return ToDo{}, nil
}

func (todo *ToDo) UpdateToDo() (result ToDo, err error) {
	return ToDo{}, nil
}

func (todo *ToDo) DeleteToDo() (err error) {
	return nil
}

func ReadToDoList() (todoList []ToDo, err error) {
	rows, err := Db.Query("SELECT Id,Subject,Priority,CreatedAt FROM ToDoList")

	if err != nil {
		log.Fatal("Connection Error : ", err)
		return
	}

	for rows.Next() {
		todo := ToDo{}

		if err = rows.Scan(&todo.Id, &todo.Subject, &todo.Priority, &todo.CreatedAt); err != nil {
			return
		}

		todoList = append(todoList, todo)
	}

	rows.Close()
	return
}
