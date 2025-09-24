package employee

import "errors"

var (
	ErrInvalidId        = errors.New("ID must be positive")
	ErrEmptyFirstName   = errors.New("first name cannot be empty")
	ErrEmptyLastName    = errors.New("last name cannot be empty")
	ErrInvalidSalaryMin = errors.New("salary must be positive")
	ErrInvalidSalaryMax = errors.New("salary mut be less than Rp.1jt")
	ErrFutureHireDate   = errors.New("hire date cannot be in the future")
)
