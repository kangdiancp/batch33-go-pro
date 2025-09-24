package employee

import (
	"strings"
	"sync/atomic"
	"time"

	"github.com/codeid/part06/pkg/department"
)

var lastID atomic.Int64

func init() {
	lastID.Store(100)
}

func GenerateID() int64 {
	return lastID.Add(1)
}

// 1.jadikan public dulu employee
type Employee struct {
	id         int64 //private
	firstName  string
	lastName   string
	hireDate   time.Time
	salary     float64
	department *department.Department
}

// 1. constructor return pointer employee
// shareing value
// encapsulation method
func NewEmployee(firstName string, lastName string, hireDate time.Time, salary float64) *Employee {
	return &Employee{
		id:        GenerateID(),
		firstName: firstName,
		lastName:  lastName,
		hireDate:  hireDate,
		salary:    salary,
	}
}

// 2. constructor return value employee
func NewEmployeeValue(firstName string, lastName string, hireDate time.Time, salary float64) Employee {
	return Employee{
		id:        GenerateID(),
		firstName: firstName,
		lastName:  lastName,
		hireDate:  hireDate,
		salary:    salary,
	}
}

// 3. constructor with department
func NewEmployeeWithDept(firstName string, lastName string, hireDate time.Time, salary float64, department *department.Department) *Employee {
	return &Employee{
		id:         GenerateID(),
		firstName:  firstName,
		lastName:   lastName,
		hireDate:   hireDate,
		salary:     salary,
		department: department,
	}
}

func NewEmployeeValid(firstName string, lastName string, hireDate time.Time, salary float64) (*Employee, error) {

	/* if salary < minimumWage {
		return nil, ErrInvalidSalaryMin
	} else if salary > maxmumWage {
		return nil, ErrInvalidSalaryMax
	}

	if strings.TrimSpace(firstName) == "" {
		return nil, ErrEmptyFirstName
	} */

	if err := validateEmployee(firstName, lastName, hireDate, salary); err != nil {
		return nil, err
	}

	return &Employee{
		id:        GenerateID(),
		firstName: firstName,
		lastName:  lastName,
		hireDate:  hireDate,
		salary:    salary,
	}, nil
}

func (e *Employee) GetId() int64 {
	if e != nil {
		return e.id
	}
	return 0
}

func (e *Employee) SetId(id int64) {
	if e != nil {
		e.id = id
	}
}

func (e *Employee) GetFirstName() string {
	if e != nil {
		return e.firstName
	}
	return ""
}

func (e *Employee) SetFirstName(firstName string) {
	if e != nil {
		e.firstName = firstName
	}
}

func (e *Employee) GetLastName() string {
	if e != nil {
		return e.lastName
	}
	return ""
}

func (e *Employee) SetLastName(lastName string) {
	if e != nil {
		e.lastName = lastName
	}
}

// TODO: default value for time.Time is not resolved.
func (e *Employee) GetHireDate() time.Time {
	var ret time.Time
	if e != nil {
		ret = e.hireDate
	}
	return ret
}

func (e *Employee) SetHireDate(hireDate time.Time) {
	if e != nil {
		e.hireDate = hireDate
	}
}

func (e *Employee) GetSalary() float64 {
	if e != nil {
		return e.salary
	}
	return 0
}

func (e *Employee) SetSalary(salary float64) error {
	if e != nil {
		if salary < minimumWage {
			return ErrInvalidSalaryMin
		} else if salary > maxmumWage {
			return ErrInvalidSalaryMax
		}
		e.salary = salary
	}
	return nil
}

func validateEmployee(firstName, lastName string, hireDate time.Time, salary float64) error {

	if strings.TrimSpace(firstName) == "" {
		return ErrEmptyLastName
	}

	if strings.TrimSpace(lastName) == "" {
		return ErrEmptyLastName
	}
	if salary < minimumWage {
		return ErrInvalidSalaryMin

	} else if salary > maxmumWage {
		return ErrInvalidSalaryMax
	}
	if hireDate.After(time.Now()) {
		return ErrFutureHireDate
	}
	return nil
}
