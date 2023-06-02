package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func (e *Employee) IncreaseSalary(amt int) {
	e.Salary += amt
}

func main() {
	manager := Employee{1, "Jake", "123 Lane st", time.Now(), "manager", 10000, 2}
	fmt.Println(manager)

	softwareEngineer := Employee{ID: 2, Name: "Engi", Salary: 342, Address: "Address", DoB: time.Now(), Position: "SWE", ManagerID: 1}
	fmt.Println(softwareEngineer)

	softwareEngineer.IncreaseSalary(100)

	fmt.Println(softwareEngineer)
}
