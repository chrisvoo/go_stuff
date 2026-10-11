package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"os"
)

type Team struct {
	Name        string
	PlayerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

type Ranker interface {
	Ranking() []string
}

func (l *League) MatchResult(firstTeamName string, firstTeamScore int, secondTeamName string, secondTeamScore int) {
	if l.Wins == nil {
		l.Wins = make(map[string]int)
	}
	if firstTeamScore > secondTeamScore {
		l.Wins[firstTeamName]++
	} else if secondTeamScore > firstTeamScore {
		l.Wins[secondTeamName]++
	}
}

func (l *League) Ranking() []string {
	teams := make([]string, 0, len(l.Teams))
	for _, team := range l.Teams {
		teams = append(teams, team.Name)
	}
	sort.Slice(teams, func(i, j int) bool {
		leftWins := l.Wins[teams[i]]
		rightWins := l.Wins[teams[j]]
		if leftWins == rightWins {
			return teams[i] < teams[j] // Alphabetical tie-breaker
		}
		return leftWins > rightWins
	})
	return teams
}

func RankPrinter(ranker Ranker, out io.Writer) {
	for _, team := range ranker.Ranking() {
		_, _ = io.WriteString(out, team)
		_, _ = io.WriteString(out, "\n")
	}
}

func main() {
	exercise1()
}

/*
You have been asked to manage a basketball league and are going to write a program to help you.
Define two types. The first one, called Team, has a field for the name of the team and a field for the player names.
The second type is called League and has a field called Teams for the teams in the league and a field called Wins
that maps a team’s name to its number of wins.
Add two methods to League.
The first method is called MatchResult. It takes four parameters: the name of the first team, its score in the game,
the name of the second team, and its score in the game. This method should update the Wins field in League.
Add a second method to League called Ranking that returns a slice of the team names in order of wins.
Build your data structures and call these methods from the main function in your program using some sample data.
Define an interface called Ranker that has a single method called Ranking that returns a slice of strings.
Write a function called RankPrinter with two parameters, the first of type Ranker and the second of type io.Writer.
Use the io.WriteString function to write the values returned by Ranker to the io.Writer, with a newline
separating each result. Call this function from main.
*/
func exercise1() {
	league := League{
		Teams: []Team{
			{Name: "Lions", PlayerNames: []string{"Ava", "Ben"}},
			{Name: "Tigers", PlayerNames: []string{"Cara", "Drew"}},
			{Name: "Bears", PlayerNames: []string{"Eli", "Fay"}},
		},
		Wins: make(map[string]int),
	}

	league.MatchResult("Lions", 82, "Tigers", 75) // Lions win
	league.MatchResult("Bears", 88, "Lions", 90)  // Lions win
	league.MatchResult("Tigers", 76, "Bears", 70) // Tigers win

	RankPrinter(&league, os.Stdout)
}

func recap() {
	nilStruct()
	constIota()
	embeddedField()
	methodPromotion()
	comparable()
	parseTree()
	injection()
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

// If you want to compare the contents of the types (e.g., checking if two slices contain [1, 2, 3]) without panicking on non-comparable types, use reflect.DeepEqual inside DoublerCompare
// func DoublerCompare(d1, d2 Doubler) {
//     fmt.Println(reflect.DeepEqual(d1, d2))
// }

/*
false: Both interface values have the same dynamic type (*DoubleInt). However, pointer comparisons in Go check memory addresses. Because di and di2 are two distinct variables residing at different memory locations, &di == &di2 evaluates to false
false: Interface equality checks require both the dynamic type and dynamic value to match. Here, the dynamic types differ (*DoubleInt vs DoubleIntSlice). Because the underlying concrete types are not identical, Go immediately evaluates the comparison to falsw
panic: runtime error: comparing uncomparable type main.DoubleIntSlice. Both interface values hold the same dynamic type (DoubleIntSlice), which is underlyingly a slice ([]int). In Go, slices are non-comparable types. While Go permits putting a slice into an interface, attempting to compare two interface values containing non-comparable dynamic types results in a runtime panic rather than a compile-time error.

While the logic inside Double() isn't involved, the receiver type in the method signature
(*DoubleInt vs DoubleIntSlice) indirectly dictates how comparison behaves because it dictates what type of value
must be stored in the interface.
*/
func comparable() {
	var di DoubleInt = 10
	var di2 DoubleInt = 10
	var dis = DoubleIntSlice{1, 2, 3}
	var dis2 = DoubleIntSlice{1, 2, 3}
	DoublerCompare(&di, &di2)
	DoublerCompare(&di, dis)
	DoublerCompare(dis, dis2)
}

/*
Inner: 20
Outer: 5
Inner: 5
*/
func methodPromotion() {
	o := Outer{
		A: 10,
		S: "Hello",
	}
	fmt.Println(o.Double())
	fmt.Println(o.IntPrinter(5))
	fmt.Println(o.Inner.IntPrinter(5))
}

func embeddedField() {
	type Employee struct {
		Name string
		ID   string
	}

	type Manager struct {
		Employee
		Reports []Employee
	}

	m := Manager{
		Name:    "Bob Bobson",
		ID:      "12345",
		Reports: []Employee{},
	}
	/* equivalent to
		 m := Manager{
	    	Employee: Employee{
				Name: "Bob Bobson",
				ID:   "12345",
			},
			Reports: []Employee{},
		 } */

	fmt.Println(m.ID)

	type Inner struct {
		X int
	}

	type Outer struct {
		Inner
		X int
	}

	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X)       // prints 20
	fmt.Println(o.Inner.X) // prints 10
}

/*
Field2 is assigned 2 because iota has a value of 1 on the second line in the const block.
Field4 is assigned 20 because it has no type or value explicitly assigned, so it gets the value of the previous line with a type and assignment.
Field5 gets the value 4 because it is the fifth line and iota starts counting from 0.
*/
func constIota() {
	const (
		Field1 = 0
		Field2 = 1 + iota
		Field3 = 20
		Field4
		Field5 = iota
	)
	fmt.Println(Field1, Field2, Field3, Field4, Field5)
}

func nilStruct() {
	var tree *IntTree
	fmt.Println("Contains method:", tree.Contains(5))

	/*
		When you call tree.Insert(5) on a nil pointer, Go passes that nil value into the method. The method
		creates a new struct and returns its memory address. However, Go passes method receivers by value
		(meaning Insert gets a copy of the nil pointer). The Insert method cannot overwrite the original
		pointer variable back in main to make it point to the new struct. That is why you must explicitly
		capture the returned memory address with tree = tree.Insert(...).
	*/
	tree = tree.Insert(5)
	fmt.Println("Contains method:", tree.Contains(5))

	leftTree := IntTree{
		val:   1,
		left:  nil,
		right: nil,
	}

	newTree := IntTree{
		val:   1,
		left:  &leftTree,
		right: nil,
	}

	/*
		newTree is an already-allocated struct value. When you call newTree.Insert(8), Go automatically takes
		the address of your struct and passes it to the pointer receiver (it acts like (&newTree).Insert(8)).
	*/
	newTree.Insert(8)
	fmt.Println("newTree:", newTree.Contains(8))
	/*
		Contains method: false
		Contains method: true
		newTree: true
	*/
}

func parseTree() {
	parseTree, err := parse("5*10+20")
	if err != nil {
		panic(err)
	}
	result, err := walkTree(parseTree)
	fmt.Println(result, err)
}

// http://localhost:8081/hello?user_id=1
func injection() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8081", nil)
}
