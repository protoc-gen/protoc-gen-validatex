package validatex

import "github.com/protoc-gen/protoc-gen-go-errors/errors"

func NewError(message string) *errors.Error {
	return errors.New(ErrCodeInvalidParameters, ErrInvalidParameters, message)
}
