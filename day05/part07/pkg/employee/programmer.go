package employee

import (
	"fmt"
	"time"

	"github.com/codeid/part07/pkg/department"
)

type Programmer struct {
	Employee
	placement string
}

func NewProgrammer(firstName string, lastName string, hireDate time.Time, salary float64, department *department.Department, placement string) *Programmer {
	return &Programmer{
		Employee:  *NewEmployeeWithDept(firstName, lastName, hireDate, salary, department),
		placement: placement,
	}
}

func (p *Programmer) GetPlacement() string {
	if p != nil {
		return p.placement
	}
	return ""
}

func (p *Programmer) SetPlacement(placement string) {
	if p != nil {
		p.placement = placement
	}
}

func (p *Programmer) ToString() string {
	return fmt.Sprintf("Id: %d, FullName: %s %s, HireDate: %s, Salary: %.2f, Placement:%s",
		p.id, p.firstName, p.lastName, p.hireDate.Format("2025-12-01"), p.salary, p.placement)
}

func (p *Programmer) ToJson() (string, error) {
	panic("not implemented") // TODO: Implement
}
