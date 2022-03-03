package e

import (
	"errors"
	"fmt"
)

type ErrInvalidParam struct {
	Param interface{}
}

func (e *ErrInvalidParam) Error() string {
	return fmt.Sprintf("invalid param: %v", e.Param)
}

func IsErrInvalidParam(e error) bool {
	var err *ErrInvalidParam
	if errors.As(e, &err) {
		return true
	}
	return false
}
