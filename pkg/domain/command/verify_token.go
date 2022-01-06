package command

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/opencars/auth/pkg/eventapi"
)

type VerifyToken struct {
	Secret string
	IP     string
}

func (c *VerifyToken) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.Secret,
			validation.Required.Error("required"),
		),
		validation.Field(
			&c.IP,
			validation.Required.Error("required"),
		),
	)
}

func (c *VerifyToken) Event() eventapi.Event {

}
