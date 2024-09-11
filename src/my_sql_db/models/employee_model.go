package models

type EmployeeModel struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	RollNo  int    `json:"rollno"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}
