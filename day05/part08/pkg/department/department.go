package department

type Department struct {
	id             int64
	departmentName string
}

func NewDepartment(id int64, departmentName string) *Department {
	return &Department{
		id:             id,
		departmentName: departmentName,
	}
}

func (d *Department) GetId() int64 {
	if d != nil {
		return d.id
	}
	return 0
}

func (d *Department) SetId(id int64) {
	if d != nil {
		d.id = id
	}
}

func (d *Department) GetDepartmentName() string {
	if d != nil {
		return d.departmentName
	}
	return ""
}

func (d *Department) SetDepartmentName(departmentName string) {
	if d != nil {
		d.departmentName = departmentName
	}
}
