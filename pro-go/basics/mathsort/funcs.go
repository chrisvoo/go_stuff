package main

import (
    "fmt"
    "math/rand"
)

func Printfln(template string, values ...interface{}) {
    fmt.Printf(template + "\n", values...)
}

func IntRange(min, max int) int {
    return rand.Intn(max - min) + min
}