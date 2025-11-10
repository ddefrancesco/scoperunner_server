package errors

import (
	"testing"
)

func TestObjectNotFoundInCatalogError_Error(t *testing.T) {
	err := ObjectNotFoundInCatalogError{
		Message: "test error message",
	}

	result := err.Error()
	if result != "test error message" {
		t.Errorf("expected 'test error message', got '%s'", result)
	}
}

func TestNewObjectNotFoundInCatalogError(t *testing.T) {
	name := "M31"
	err := NewObjectNotFoundInCatalogError(name)

	expectedMessage := "NGC object with Name M31 not found"
	if err.Message != expectedMessage {
		t.Errorf("expected '%s', got '%s'", expectedMessage, err.Message)
	}

	// Test that it implements error interface
	var _ error = err
	
	// Test Error() method
	errorString := err.Error()
	if errorString != expectedMessage {
		t.Errorf("Error() method: expected '%s', got '%s'", expectedMessage, errorString)
	}
}