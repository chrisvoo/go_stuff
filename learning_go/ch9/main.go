package main

import (
	"ch9/exercises"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func main() {
	exe1()
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

/*
Create a sentinel error to represent an invalid ID. In main, use errors.Is to check
for the sentinel error, and print a message when it is found.
*/
func exe1() {
	// json.NewDecoder operates on streams, not single documents.
	// it does not try to parse the entire input at once; instead, it parses values
	// sequentially, one at a time, moving an internal cursor through the stream.
	d := json.NewDecoder(strings.NewReader(data))
	count := 0

	/*
		Under the hood in Go's encoding/json, dec.More() works essentially like this

		func (dec *Decoder) More() bool {
			c, err := dec.peek()
			return err == nil && c != ']' && c != '}'
		}

		The following code is more idiomatic for reading top-level concatenated JSON streams

		for {
			var emp exercise1.Employee
			err := d.Decode(&emp)
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				// handle syntax or decoding error
			}
			// process emp
		}
	*/

	for d.More() {
		count++
		var emp exercises.Employee

		// 1. It skips any leading whitespace (newlines, spaces, tabs).
		// 2. It reads tokens until it finds the balanced closing brace } for that object.
		// 3. It unmarshals that single object into the emp struct and pauses reader consumption right after the closing brace.
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = exercises.ValidateEmployeeMultipleErrors(emp)
		message := fmt.Sprintf("record %d: %+v", count, emp)

		if err != nil {
/* 			handleSingleError(err, count, emp)
			continue */
			message = handleMultipleErrors(err, message, emp)
		}
		fmt.Println(message)
	}
}

/*
Define a custom error type to represent an empty field error. This error should include the
name of the empty Employee field. In main, use errors.As to check for this error.
Print out a message that includes the field name.
*/
func handleSingleError(err error, count int, emp exercises.Employee) {
	// define a variable of type EmptyFieldError to use with errors.As
	var fieldErr exercises.EmptyFieldError

	// Use errors.Is to check for the sentinel error.
	if errors.Is(err, exercises.ErrInvalidID) {
		fmt.Printf("record %d: %+v error: invalid ID: %s\n", count, emp, emp.ID)
	} else if errors.As(err, &fieldErr) {
		// use errors.As to check if the error returned is (or wraps) an EmptyFieldError
		fmt.Printf("record %d: %+v error: empty field %s\n", count, emp, fieldErr.FieldName)
	} else {
		fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
	}
}

func handleMultipleErrors(err error, message string, emp exercises.Employee) string {
	// use a type switch to detect an error that can be unwrapped into multiple errors
	switch err := err.(type) {
	case interface{ Unwrap() []error }:
		// if we have one use a loop to walk through allErrors and build a single string
		allErrors := err.Unwrap()
		var messages []string
		for _, e := range allErrors {
			messages = append(messages, exercises.ProcessError(e, emp))
		}
		message = message + " allErrors: " + strings.Join(messages, ", ")
	default:
		message = message + " error: " + exercises.ProcessError(err, emp)
	}

	return message
}
