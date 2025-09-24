package main

import (
	"fmt"
	"time"

	"github.com/codeid/part07/pkg/department"
	"github.com/codeid/part07/pkg/employee"
)

func main() {

	//1.initial value deparatment
	deptIt := department.NewDepartment(1, "IT")
	deptFinance := department.NewDepartment(2, "FINANCE")
	deptSales := department.NewDepartment(3, "SALES")

	//1. constructor return pointer employee (recommended)
	emp1 := employee.NewEmployeeWithDept("Kang", "Dian", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 500_000, deptIt)
	emp2 := employee.NewProgrammer("Rini", "Maharani", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 500_000, deptFinance, "BCA")
	emp3 := employee.NewProgrammer("Budi", "Juna", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 250_000, deptIt, "IT")
	emp4 := employee.NewProgrammer("Charlie", "Chaplin", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 450_000, deptFinance, "BRI")
	emp5 := employee.NewManager("Dadan", "Galon", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 55_000, deptSales, 10)

	//2. ToString() hanya bisa menampilkan data employee, informasi tambahan seperti placement & totalStaff ga muncul
	employees := []*employee.Employee{emp1, &emp2.Employee, &emp3.Employee, &emp4.Employee, &emp5.Employee}

	for _, emp := range employees {
		fmt.Println(emp.ToString())
	}

	//3. poopulate with interface
	employeesInf := []employee.Info{emp1, emp2, emp3, emp4, emp5}
	for _, v := range employeesInf {
		fmt.Println(v.ToString())
	}

	/* 	fmt.Println(emp1.ToString())
	   	fmt.Println(emp2.ToString())
	   	fmt.Println(emp3.ToString())
	   	fmt.Println(emp4.ToString())
	   	fmt.Println(emp5.ToString()) */

}
