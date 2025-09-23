package main

import (
	"fmt"
	"time"
)

type employee struct {
	firstName string
	lastName  string
	hireDate  time.Time
	salary    float64
}

func toString(e employee) string {
	return fmt.Sprintf("FullName: %s %s, HireDate: %s, Salary: %.2f",
		e.firstName, e.lastName, e.hireDate.Format("2025-02-02"), e.salary)
}

func main() {

	//1. literal struct
	emp1 := employee{firstName: "Kang", lastName: "Dian", hireDate: time.Now(), salary: 100_000}
	//fmt.Printf("emp1:%v\n", emp1)
	fmt.Println(toString(emp1))

	//2. urutan field
	emp2 := employee{"Budi", "Juna", time.Now(), 150_000}
	//fmt.Printf("emp2:%v\n", emp2)
	fmt.Println(toString(emp2))

	//3. using constructor
	emp3 := new(employee)
	emp3.firstName = "Charlie"
	emp3.lastName = "Chaplin"
	emp3.hireDate = time.Now()
	emp3.salary = 125_00

	fmt.Println(toString(*emp3))

	//fmt.Printf("emp3:%v\n", emp3)

}
