package employee

import (
	"encoding/json"
	"fmt"
)

func (m *Manager) ToString() string {
	return fmt.Sprintf("Id: %d, FullName: %s %s, HireDate: %s, Salary: %.2f, TotalStaff:%d",
		m.id, m.firstName, m.lastName, m.hireDate.Format("2025-12-01"), m.salary, m.totalStaff)
}

func (m *Manager) ToJson() (string, error) {
	data := map[string]any{
		"id":         m.id,
		"firstName":  m.firstName,
		"lastName":   m.lastName,
		"hireDate":   m.hireDate.Format("2025-12-01"),
		"salary":     m.salary,
		"totalStaff": m.totalStaff,
	}
	jsonBytes, err := json.MarshalIndent(data, "", " ")
	return string(jsonBytes), err
}
