package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mydatabase/models"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

func insertOneEmployee(movie models.EmployeeModel) error {
	query := "INSERT INTO employee(name,rollno,age,address) VALUES(?,?,?,?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	fmt.Println(movie.Name, movie.Age, movie.Address)

	res, err := stmt.ExecContext(ctx, movie.Name, movie.RollNo, movie.Age, movie.Address)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created ", rows)
	return nil
}

func getOneEmployee(id string) []models.EmployeeModel {
	res, err := db.Query("select * from employee where id= ?", id)
	_, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
	}
	fmt.Println("near model")
	var employees []models.EmployeeModel
	for res.Next() {
		var sEmployee models.EmployeeModel
		if err := res.Scan(&sEmployee.ID, &sEmployee.Name, &sEmployee.RollNo, &sEmployee.Age, &sEmployee.Address); err != nil {
			fmt.Println(err)
			return employees
		}
		employees = append(employees, sEmployee)
	}
	return employees
}

func getAllEmployees() []models.EmployeeModel {
	res, err := db.Query("select * from employee")
	_, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
	}
	var employees []models.EmployeeModel
	for res.Next() {
		var sEmployee models.EmployeeModel
		if err := res.Scan(&sEmployee.ID, &sEmployee.Name, &sEmployee.RollNo, &sEmployee.Age, &sEmployee.Address); err != nil {
			fmt.Println(err)
			return employees
		}
		employees = append(employees, sEmployee)
	}
	return employees
}

func deleteOneEmployee(id string) {
	query := "DELETE FROM employee WHERE id = ?"
	res, err := db.Exec(query, id)
	_, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.LastInsertId())
}

func updateOneEmployee(id string, emp models.EmployeeModel) string {
	query := "UPDATE employee SET name = ?, age = ?, address = ?, rollno = ? WHERE id = ?"
	result, err := db.Exec(query, emp.Name, emp.Age, emp.Address, emp.RollNo, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())
	return id
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	myemployee := getAllEmployees()
	json.NewEncoder(w).Encode(myemployee)
}
func GetOneEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	myemployee := getOneEmployee(params["id"])
	json.NewEncoder(w).Encode(myemployee)
}
func CreateOneEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.EmployeeModel
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Fatal(err)
	}
	insertOneEmployee(employee)
	json.NewEncoder(w).Encode(employee)
}

func DeleteOneEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	deleteOneEmployee(id)
	json.NewEncoder(w).Encode("The emplee deleted successfully" + id)
}

func UpdateOneEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var emp models.EmployeeModel
	json.NewDecoder(r.Body).Decode(&emp)
	var empId = updateOneEmployee(params["id"], emp)
	json.NewEncoder(w).Encode("employee updated " + empId)
}

// task creation
// 1) task id, task name, task description, task status, task start date, task end date

//task fetch by id

//fetch all tasks

//change task status

//delete tasks

//edit task
