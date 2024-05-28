package http_request

import "github.com/go-playground/validator/v10"

var (
	Val = validator.New()
)

func Validate(v interface{}) error {
	return Val.Struct(v)
}
