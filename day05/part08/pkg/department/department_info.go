package department

import (
	"encoding/json"
	"fmt"
)

func (d *Department) ToString() string {
	return fmt.Sprintf("ID : %d, DepartmentName : %s", d.id, d.departmentName)
}

func (d *Department) ToJson() (string, error) {
	data := map[string]any{
		"id":             d.id,
		"departmentName": d.departmentName,
	}
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	return string(jsonBytes), err
}
