package employee

import (
	"encoding/json"
	"fmt"
)

func (m *Programmer) ToString() string {
	return fmt.Sprintf("Id: %d, FullName: %s %s, HireDate: %s, Salary: %.2f, Placement:%s",
		m.id, m.firstName, m.lastName, m.hireDate.Format("2025-12-01"), m.salary, m.Placement)
}

func (m *Programmer) ToJson() (string, error) {
	data := map[string]any{
		"id":        m.id,
		"firstName": m.firstName,
		"lastName":  m.lastName,
		"hireDate":  m.hireDate.Format("2025-12-01"),
		"salary":    m.salary,
		"placement": m.Placement,
	}
	jsonBytes, err := json.MarshalIndent(data, "", " ")
	return string(jsonBytes), err
}
