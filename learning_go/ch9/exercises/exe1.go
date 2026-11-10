package exercises

import (
	"errors"
	"fmt"
	"regexp"
)

var ErrInvalidID = errors.New("invalid ID")

type EmptyFieldError struct {
	FieldName string
}

/*
Define an Error() string method on EmptyFieldError so it meets the error interface.
*/
func (fe EmptyFieldError) Error() string {
	return fe.FieldName
}

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

func ValidateEmployee(e Employee) error {
	if len(e.ID) == 0 {
		return EmptyFieldError{FieldName: "ID"}
	}
	if !validID.MatchString(e.ID) {
		// Change the error returned to ErrInvalidID.
		return ErrInvalidID
	}
	if len(e.FirstName) == 0 {
		return EmptyFieldError{FieldName: "FirstName"}
	}
	if len(e.LastName) == 0 {
		return EmptyFieldError{FieldName: "LastName"}
	}
	if len(e.Title) == 0 {
		return EmptyFieldError{FieldName: "missing Title"}
	}
	return nil
}

func ValidateEmployeeMultipleErrors(e Employee) error {
	// use allErrors to gather the errors
	var allErrors []error
	if len(e.ID) == 0 {
		// append on to allErrors
		allErrors = append(allErrors, EmptyFieldError{FieldName: "ID"})
	} else if !validID.MatchString(e.ID) {
		// append on to allErrors
		allErrors = append(allErrors, ErrInvalidID)
	}
	if len(e.FirstName) == 0 {
		// append on to allErrors
		allErrors = append(allErrors, EmptyFieldError{FieldName: "FirstName"})
	}
	if len(e.LastName) == 0 {
		// append on to allErrors
		allErrors = append(allErrors, EmptyFieldError{FieldName: "LastName"})
	}
	if len(e.Title) == 0 {
		// append on to allErrors
		allErrors = append(allErrors, EmptyFieldError{FieldName: "Title"})
	}
	// depending on the number of errors found, return nil, the single error, or use
	// errors.Join to return all errors within a single error
	switch len(allErrors) {
	case 0:
		return nil
	case 1:
		return allErrors[0]
	default:
		return errors.Join(allErrors...)
	}
}

/*
processError is a method for converting an error into a message. It has special handling for ErrInvalidID and
EmptyFieldError.
*/
func ProcessError(err error, emp Employee) string {
	var fieldErr EmptyFieldError
	if errors.Is(err, ErrInvalidID) {
		return fmt.Sprintf("invalid ID: %s", emp.ID)
	} else if errors.As(err, &fieldErr) {
		return fmt.Sprintf("empty field %s", fieldErr.FieldName)
	} else {
		return fmt.Sprintf("%v", err)
	}
}
