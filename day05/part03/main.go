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

func GetFullInfo(e employee) string {

	return fmt.Sprintf("FullName: %s %s, HireDate:%s, Salary: %.2f, Department: %s, Location:%s",
		e.firstName, e.lastName, e.hireDate.Format("2006-01-02"), e.salary, e.department.departmentName, e.department.location.address)
}

func main() {
	//4. embedded struct
	emp1 := employee{
		firstName: "kang",
		lastName:  "dian",
		hireDate:  time.Now(),
		salary:    500_000,
		department: department{
			id:             1,
			departmentName: "IT",
			location:       location{id: 1, address: "jakarta"},
		},
	}

	//fmt.Printf("emp1 : %v\n", emp1)
	fmt.Println(GetFullInfo(emp1))

	emp1.firstName = "Reni"
	emp1.lastName = "Mereni"
	emp1.salary = 550_000

	//fields promotions
	emp1.department.departmentName = "Finance"
	emp1.department.location.address = "Bandung"

	fmt.Println(GetFullInfo(emp1))

}
