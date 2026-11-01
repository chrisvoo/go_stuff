package main

import (
	"ch8/stack"
	"ch8/tree"
	"ch8/exercises"
	"ch8/fmr"
	"cmp"
	"fmt"
)

func main() {
	exercises.Exercise1()
	exercises.Exercise2()
	exercises.Exercise3()
}

func recap() {
	stackFunc()
	filters()
	treeGenerics()
}

func stackFunc() {
	// var intStack Stack[int]
	s := stack.Stack[int]{Vals: make([]int, 0)}
	s.Push(5)
	s.Push(7)
	s.Push(9)
	val, _ := s.Pop()
	fmt.Printf("Removed %d\n", val)
	fmt.Println("Contains 9?", s.Contains(9))
	fmt.Printf("Total elements: %d\n", len(s.Vals))
}



func filters() {
	words := []string{"One", "Potato", "Two", "Potato"}
	filtered := fmr.Filter(words, func(s string) bool {
		return s != "Potato"
	})
	fmt.Println(filtered)
	lengths := fmr.Map(filtered, func(s string) int {
		return len(s)
	})
	fmt.Println(lengths)
	sum := fmr.Reduce(lengths, 0, func(acc int, val int) int {
		return acc + val
	})
	fmt.Println(sum)
}

/*
true
false
true
false
true
false
*/
func treeGenerics() {
	t1 := tree.NewTree(cmp.Compare[int])
	t1.Add(10)
	t1.Add(30)
	t1.Add(15)
	fmt.Println(t1.Contains(15))
	fmt.Println(t1.Contains(40))

	t2 := tree.NewTree(tree.OrderPeople)
	t2.Add(tree.Person{Name: "Bob", Age: 30})
	t2.Add(tree.Person{Name: "Maria", Age: 35})
	t2.Add(tree.Person{Name: "Bob", Age: 50})
	fmt.Println(t2.Contains(tree.Person{Name: "Bob", Age: 30}))
	fmt.Println(t2.Contains(tree.Person{Name: "Fred", Age: 25}))

	t3 := tree.NewTree(tree.Person.Order)
	t3.Add(tree.Person{Name: "Bob", Age: 30})
	t3.Add(tree.Person{Name: "Maria", Age: 35})
	t3.Add(tree.Person{Name: "Bob", Age: 50})
	fmt.Println(t3.Contains(tree.Person{Name: "Bob", Age: 30}))
	fmt.Println(t3.Contains(tree.Person{Name: "Fred", Age: 25}))
}
