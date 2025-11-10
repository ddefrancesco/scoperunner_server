package errors

import "fmt"

type ObjectNotFoundInCatalogError struct {
	Message string
}

func (e ObjectNotFoundInCatalogError) Error() string {
	return e.Message
}

func NewObjectNotFoundInCatalogError(name string) ObjectNotFoundInCatalogError {
	return ObjectNotFoundInCatalogError{
		Message: fmt.Sprintf("NGC object with Name %v not found", name),
	}
}
