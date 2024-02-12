package handler

import "fmt"

func errParamRequired(name string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {

	if r.Role == "" || r.Company == "" || r.Location == "" || r.Remote == nil || r.Salary <= 0 {
		return fmt.Errorf("request body is empty or contains invalid values!")
	}
	if r.Role == "" {
		return errParamRequired("role", "string")
	}
	if r.Company == "" {
		return errParamRequired("company", "string")
	}
	if r.Location == "" {
		return errParamRequired("location", "string")
	}
	if r.Remote == nil {
		return errParamRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamRequired("salary", "int64")
	}
	return nil
}
