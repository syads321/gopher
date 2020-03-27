package types

import (
	"fmt"
)

type CommonError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e CommonError) Error() string {
	return fmt.Sprintf("error [%s]: %s", e.Code, e.Message)
}

func (e CommonError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}
