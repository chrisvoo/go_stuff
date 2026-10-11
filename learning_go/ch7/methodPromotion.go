package main

import "fmt"

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

// When you embed a struct in Go, all the fields and methods of the inner struct are automatically promoted to the outer struct.
// Because of this promotion, Go allows you to call o.Double() directly. Under the hood,
// the Go compiler simply translates your call into o.Inner.Double()
type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}