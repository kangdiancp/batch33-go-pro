package employee

import (
	"time"

	"github.com/codeid/part08/pkg/department"
)

type Programmer struct {
	Employee
	Placement Placement
}

func NewProgrammer(firstName string, lastName string, hireDate time.Time, salary float64, department *department.Department, placement Placement) *Programmer {
	return &Programmer{
		Employee:  *NewEmployeeWithDept(firstName, lastName, hireDate, salary, department),
		Placement: placement,
	}
}

func (p *Programmer) GetPlacement() Placement {
	if p != nil {
		return p.Placement
	}
	return ""
}

func (p *Programmer) SetPlacement(placement Placement) {
	if p != nil {
		p.Placement = placement
	}
}
