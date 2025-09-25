package employee

import (
	"encoding/json"
	"fmt"
)

func (e *Employee) ToString() string {
	return fmt.Sprintf("Id: %d, FullName: %s %s, HireDate: %s, Salary: %.2f",
		e.id, e.firstName, e.lastName, e.hireDate.Format("2006-01-02"), e.salary)
}

func (e *Employee) ToJson() (string, error) {
	data := map[string]any{
		"id":        e.id,
		"firstName": e.firstName,
		"lastName":  e.lastName,
		"hireDate":  e.hireDate.Format("2025-12-01"),
		"salary":    e.salary,
	}
	jsonBytes, err := json.MarshalIndent(data, "", " ")
	return string(jsonBytes), err

}
