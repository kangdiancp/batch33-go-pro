package main

import (
	"fmt"
	"time"
)

type location struct {
	id      int64
	address string
}

type department struct {
	id             int64
	departmentName string
	location       location
}

type employee struct {
	firstName  string
	lastName   string
	hireDate   time.Time
	salary     float64
	department department
}

func toString(e employee) string {
	return fmt.Sprintf("FullName: %s %s, HireDate: %s, Salary: %.2f, Department:%s, Address:%s",
		e.firstName, e.lastName, e.hireDate.Format("2025-02-02"), e.salary,
		e.department.departmentName, e.department.location.address)
}

func main() {

	//1.literal
	emp1 := employee{
		firstName: "Kang",
		lastName:  "Dian",
		hireDate:  time.Now(),
		salary:    150_000,
		department: department{
			id:             10,
			departmentName: "IT",
			location: location{
				id:      1001,
				address: "Jakarta",
			},
		},
	}

	fmt.Println(toString(emp1))

}
