package main

import (
	"fmt"
	"time"

	"github.com/codeid/part06/pkg/department"
	"github.com/codeid/part06/pkg/employee"
)

func main() {

	//1.initial value deparatment
	deptIt := department.NewDepartment(1, "IT")
	deptFinance := department.NewDepartment(2, "FINANCE")
	deptSales := department.NewDepartment(3, "SALES")

	//1. constructor return pointer employee (recommended)
	emp1 := employee.NewEmployeeWithDept("Kang", "Dian", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 500_000, deptIt)
	emp2 := employee.NewEmployeeWithDept("Rini", "Maharani", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 500_000, deptIt)
	emp3 := employee.NewEmployeeWithDept("Budi", "Juna", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 250_000, deptFinance)
	emp4 := employee.NewEmployeeWithDept("Charlie", "Chaplin", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 450_000, deptFinance)
	emp5 := employee.NewEmployeeWithDept("Dadan", "Galon", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 55_000, deptSales)

	employees := []*employee.Employee{emp1, emp2, emp3, emp4, emp5}

	for i, v := range employees {
		fmt.Printf("Emp[%d] adress[%p] value[%v]\n", i, v, v)
	}

}
