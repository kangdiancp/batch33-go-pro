package main

import (
	"fmt"
	"time"
)

func main() {

	//1. declare variable with data type
	var fullName string = "Jhone Doe"
	fmt.Println(fullName)

	//2. declare multiple variable with datatype using var
	var (
		firstName string    = "Jhone"
		lastName  string    = "Snow"
		salary    float64   = 5_000_000
		hireDate  time.Time = time.Now()
	)

	fmt.Println(firstName, lastName, salary, hireDate)

	//3. declare multiple variable without datatype
	var (
		departId       = 10
		departmentName = "Finance"
	)

	fmt.Println(departId, departmentName)

	//4 declare mulitple variable in one-line, with zero value
	var salary1, salary2, salary3 float64
	var myName string
	fmt.Println(salary1, salary2, salary3)
	fmt.Println(myName)

	salary1 = 98.00
	fmt.Println(salary1)

	//5. short variable
	salary4 := 12.00
	tax := 0.5
	fmt.Println(salary4, tax)

	//6. multiple variable with shortvariable
	isActive, customerName, orderDate, price := false, "PT. Code", time.Now(), 550_00
	fmt.Println(isActive, customerName, orderDate, price)

	//7. constanta
	const pajak = 100
	fmt.Println(pajak)
}
