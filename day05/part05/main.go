package main

import (
	"fmt"
	"time"

	"github.com/codeid/part05/pkg/employee"
)

func main() {

	//1. constructor return pointer employee (recommended)
	emp1 := employee.NewEmployee("Kang", "Dian", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 500_000)
	emp2 := employee.NewEmployee("Rini", "Maharani", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 500_000)
	emp3 := employee.NewEmployee("Budi", "Juna", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 250_000)
	emp4 := employee.NewEmployee("Charlie", "Chaplin", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 450_000)
	emp5 := employee.NewEmployee("Dadan", "Galon", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 55_000)

	//2. append to slice pointer employee, performance, lebih hemat resource.
	employeesPtr := []*employee.Employee{emp1, emp2, emp3, emp4, emp5}

	for i, v := range employeesPtr {
		fmt.Printf("Emp[%d] adress[%p] value[%v]\n", i, v, v)
	}

	//3. update value object employee[0]
	emp1Update := employeesPtr[0]
	emp1Update.SetFirstName("Nana")
	emp1Update.SetSalary(430_000)

	fmt.Printf("Emp[%d] adress[%p] value[%v]\n", 0, emp1, emp1)

	//4.using slice employee value
	employeesValue := []employee.Employee{*emp1, *emp2, *emp3, *emp4, *emp5}

	for i, v := range employeesValue {
		fmt.Printf("Emp[%d] adress[%p] value[%v]\n", i, &v, v)
	}

	//5. update emp1
	emp2Update := employeesValue[2]
	emp2Update.SetFirstName("Chef")
	emp2Update.SetSalary(150_000)

	fmt.Printf("Emp[%d] adress[%p] value[%v]\n", 2, emp3, emp3)

}
