package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"to-do-list/models"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "rakesh12",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "employee",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
	}
	fmt.Println("connected")
}

func createTask(task models.Task) {
	query := "insert into tasklist (taskname,taskdescription,taskstatus,startdate,enddate) values(?,?,?,?,?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Exec(task.TaskName, task.TaskDescription, task.TaskStatus, task.TaskStartDate, task.TaskEndDate)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows.RowsAffected())
}

func changeTaskStatus(status string, id string) {
	query := "update tasklist set taskstatus= ? where id = ?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Exec(status, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows.RowsAffected())
}

func deleteTask(id string) {
	res, err := db.Exec("delete from tasklist where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.RowsAffected())
}

func ChangeTaskStatus(w http.ResponseWriter, r *http.Request) {
	urlstr := r.URL
	var status = urlstr.Query().Get("status")
	var id = urlstr.Query().Get("id")
	changeTaskStatus(status, id)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var myTask models.Task
	_ = json.NewDecoder(r.Body).Decode(&myTask)
	createTask(myTask)
	json.NewEncoder(w).Encode(myTask)

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	urlstr := r.URL
	var id = urlstr.Query().Get("id")
	deleteTask(id)
}
