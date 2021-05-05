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

func (todo *ToDo) CreateToDo() (lastinsertid int64, err error) {
	stmt, err := Db.Prepare("INSERT INTO todolist(Subject,Priority,CreatedAt)VALUES(?,?,NOW())")

	if err != nil {
		log.Fatal("Insert Prepare error : ", err)
		return
	}

	defer stmt.Close()

	ret, err := stmt.Exec(todo.Subject, todo.Priority)

	if err != nil {
		log.Fatal("Insert Exec error : ", err)
		return
	}

	id, _ := ret.LastInsertId()

	return id, nil
}

func (todo *ToDo) UpdateToDo() (err error) {
	stmt, err := Db.Prepare("UPDATE todolist SET Subject = ?,Priority = ? WHERE Id = ?")

	if err != nil {
		log.Fatal("Insert Prepare error : ", err)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(todo.Subject, todo.Priority, todo.Id)

	if err != nil {
		log.Fatal("Insert Exec error : ", err)
		return
	}

	return
}

func (todo *ToDo) DeleteToDo() (err error) {
	stmt, err := Db.Prepare("DELETE FROM  todolist WHERE Id = ?")

	if err != nil {
		log.Fatal("Insert Prepare error : ", err)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(todo.Id)

	if err != nil {
		log.Fatal("Insert Exec error : ", err)
		return
	}

	return
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

func ReadToDo(id int) (todo ToDo, err error) {
	todo = ToDo{}

	err = Db.QueryRow("SELECT Id,Subject,Priority,CreatedAt FROM ToDoList WHERE Id = ?", id).
		Scan(&todo.Id, &todo.Subject, &todo.Priority, &todo.CreatedAt)

	if err != nil {
		log.Fatal("Connection Error : ", err)
		return
	}

	return
}
