package main

import (
	"fmt"
	"time"
)

type employee struct {
	id        int64
	firstName string
	lastName  string
	hireDate  time.Time
	salary    float64
}

type manager struct {
	employee
	email      string
	totalStaff int64
}

type programmer struct {
	employee
	placement string
}

func toString(e employee) string {
	return fmt.Sprintf("Id: %d, FullName: %s %s, HireDate: %s, Salary: %.2f",
		e.id, e.firstName, e.lastName, e.hireDate.Format("2006-01-02"), e.salary)
}

func toStringMgr(e manager) string {
	return fmt.Sprintf("Id: %d, FullName: %s %s, HireDate: %s, Salary: %.2f Email:%s, TotalStaff:%d",
		e.id, e.firstName, e.lastName, e.hireDate.Format("2006-01-02"), e.salary, e.email, e.totalStaff)
}

func toStringProg(e programmer) string {
	return fmt.Sprintf("Id: %d, FullName: %s %s, HireDate: %s, Salary: %.2f Placement:%s",
		e.id, e.firstName, e.lastName, e.hireDate.Format("2006-01-02"), e.salary, e.placement)
}

func main() {
	//1. create object manager
	mgr1 := manager{
		employee: employee{
			firstName: "Alicia",
			lastName:  "Juned",
			hireDate:  time.Now(),
			salary:    200_000,
		},
		email:      "alicia@code.id",
		totalStaff: 10,
	}

	//2. obejct programmer
	prog1 := programmer{
		employee: employee{
			firstName: "Yudit",
			lastName:  "Elanda",
			hireDate:  time.Now(),
			salary:    250_000,
		},
		placement: "internal project",
	}

	fmt.Println(toString(mgr1.employee))
	fmt.Println(toString(prog1.employee))
	fmt.Println(toStringMgr(mgr1))
	fmt.Println(toStringProg(prog1))

}
