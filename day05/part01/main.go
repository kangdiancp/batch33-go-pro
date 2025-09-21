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

func main() {

	//1. literal struct
	emp1 := employee{firstName: "Kang", lastName: "Dian", hireDate: time.Now(), salary: 500_000}
	fmt.Printf("emp1 : %v\n", emp1)

	//2. sesuaikan dengan urutan fields
	emp2 := employee{"Budi", "Juned", time.Now(), 450_000}
	fmt.Printf("emp2 : %v\n", emp2)

	//3. constructor
	emp3 := new(employee)
	emp3.firstName = "Rini"
	emp3.lastName = "Widi"
	emp3.hireDate = time.Now()
	emp3.salary = 150_000
	fmt.Printf("emp3 : %v\n", emp3)

}
