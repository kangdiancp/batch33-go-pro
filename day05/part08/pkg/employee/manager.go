package employee

import (
	"time"

	"github.com/codeid/part08/pkg/department"
)

type Manager struct {
	Employee
	totalStaff int64
}

func NewManager(firstName string, lastName string, hireDate time.Time, salary float64, department *department.Department, totalStaff int64) *Manager {
	return &Manager{
		Employee:   *NewEmployeeWithDept(firstName, lastName, hireDate, salary, department),
		totalStaff: totalStaff,
	}
}

func (m *Manager) GetTotalStaff() int64 {
	if m != nil {
		return m.totalStaff
	}
	return 0
}

func (m *Manager) SetTotalStaff(totalStaff int64) {
	if m != nil {
		m.totalStaff = totalStaff
	}
}

/* func (m *Manager) ToString() string {
	return fmt.Sprintf("Id: %d, FullName: %s %s, HireDate: %s, Salary: %.2f, TotalStaff:%d",
		m.id, m.firstName, m.lastName, m.hireDate.Format("2025-12-01"), m.salary, m.totalStaff)
}

func (m *Manager) ToJson() (string, error) {
	panic("not implemented") // TODO: Implement
} */
