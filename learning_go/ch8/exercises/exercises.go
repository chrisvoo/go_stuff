package exercises

import (
	"fmt"
	"strconv"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func doubleNumber[T Number](num T) T {
	return num * 2
}

/* Write a generic function that doubles the value of any integer or float that’s passed in to it. Define any needed generic interfaces. */
func Exercise1() {
	var result any

	result = doubleNumber(5)
	fmt.Println("Invokin doubleNumber on 5: ", result)

	result = doubleNumber(5.5)
	fmt.Println("Invokin doubleNumber on 5.5: ", result)
}

/*
Define a generic interface called Printable that matches a type that implements fmt.Stringer
and has an underlying type of int or float64.
Define types that meet this interface. Write a function that takes in a Printable and
prints its value to the screen using fmt.Println.
*/
func Exercise2() {
	var i PrintInt = 20
	PrintIt(i)

	var f PrintFloat = 10.23
	PrintIt(f)
}

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintInt int

func (pi PrintInt) String() string {
	return strconv.Itoa(int(pi))
}

type PrintFloat float64

func (pf PrintFloat) String() string {
	return fmt.Sprintf("%f", pf)
}

func PrintIt[T Printable](t T) {
	fmt.Println(t)
}

/*
Write a generic singly linked list data type. Each element can hold a comparable value and
has a pointer to the next element in the list.
The methods to implement are as follows:
adds a new element to the end of the linked list
Add(T)
adds an element at the specified position in the linked list
Insert(T, int)
returns the position of the supplied value, -1 if it's not present
Index (T) int
*/
func Exercise3() {
	l := &List[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))

	l.Insert(100, 0)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	l.Insert(200, 1)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(200))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(300, 10)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Add(400)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(500, 6)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}
}

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *List[T]) Add(t T) {
	n := &Node[T]{
		Value: t,
	}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = l.Tail.Next
}

func (l *List[T]) Insert(t T, pos int) {
	n := &Node[T]{
		Value: t,
	}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	if pos <= 0 {
		n.Next = l.Head
		l.Head = n
		return
	}
	curNode := l.Head
	for i := 1; i < pos; i++ {
		if curNode.Next == nil {
			curNode.Next = n
			l.Tail = curNode.Next
			return
		}
		curNode = curNode.Next
	}
	n.Next = curNode.Next
	curNode.Next = n
	if l.Tail == curNode {
		l.Tail = n
	}
}

func (l *List[T]) Index(t T) int {
	i := 0
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		if curNode.Value == t {
			return i
		}
		i++
	}
	return -1
}