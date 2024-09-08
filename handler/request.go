package handler

import (
	"fmt"
	"net/url"

	"github.com/DiogoJorge1401/golang-api/schemas"
)

type CreateOpening struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int    `json:"salary"`
}

func errParamIsRequired(name, tpe string) error {
	return fmt.Errorf("param: %v (type: %v) is required", name, tpe)
}

func (c *CreateOpening) Validate() error {
	if c.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if c.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if c.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if c.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if c.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if _, err := url.ParseRequestURI(c.Link); err != nil {
		return fmt.Errorf("param: link (type: string) is not a valid URL")
	}
	if c.Salary <= 0 {
		return fmt.Errorf("param: salary (type: int) must be a positive number")
	}

	return nil
}

func (c *CreateOpening) GetOpening() *schemas.Opening {
	return &schemas.Opening{
		Role:     c.Role,
		Company:  c.Company,
		Location: c.Location,
		Remote:   *c.Remote,
		Link:     c.Link,
		Salary:   c.Salary,
	}
}
