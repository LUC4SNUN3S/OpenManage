package handler

import "fmt"

func paramIsRequired(param, typ string) error {
	return fmt.Errorf("param %s (type: %s) is required", param, typ)
}

type OpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

type CreateOpeningRequest struct {
	OpeningRequest
}

func (opening *CreateOpeningRequest) Validate() error {
	if opening.Role == "" && opening.Company == "" && opening.Location == "" && opening.Remote == nil && opening.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if opening.Role == "" {
		return paramIsRequired("role", "string")
	}

	if opening.Company == "" {
		return paramIsRequired("company", "string")
	}

	if opening.Location == "" {
		return paramIsRequired("location", "string")
	}

	if opening.Link == "" {
		return paramIsRequired("link", "string")
	}

	if opening.Salary <= 0 {
		return paramIsRequired("salary", "int64")
	}

	if opening.Remote == nil {
		return paramIsRequired("remote", "bool")
	}

	return nil
}

type UpdateOpeningRequest struct {
	OpeningRequest
}

func (opening *UpdateOpeningRequest) Validate() error {
	if opening.Role == "" || opening.Company == "" || opening.Location == "" || opening.Remote == nil || opening.Salary <= 0 {
		return nil
	}

	return fmt.Errorf("request body is empty or malformed")
}
