package utils

// A place for REAL useful helper functions.

import (
	"arman-estimation-service/types/structs"
	"fmt"
)

func CustomError(category structs.ErrorCategory, err string, bindings ...any) *structs.CustomError {
	return &structs.CustomError{Err: fmt.Sprintf(err, bindings...), Category: category}
}
