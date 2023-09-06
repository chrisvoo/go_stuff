package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type TestStruct struct {
	Sample    string `json:"sample"`
	SampleInt int    `json:"sample_int"`
}

/*
 * If we wanted to manipulate the initialized struct itself and not the copy of the struct,
 * we will need to add an asterisk to the definition of the method.
 */
func (t *TestStruct) ChangeSample(s string) {
	t.Sample = s
}

func prettyPrint(t TestStruct) {
	jsonData, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", jsonData)
}

func newTopic(topic string) {
	fmt.Printf("%s -----------------------\n", topic)
}

func endTopic() {
	fmt.Printf("\n")
}

// A main function executes by default when you run the main package.
func main() {
	newTopic("Struct")
	testVar := TestStruct{Sample: "aa", SampleInt: 1}
	fmt.Printf("%+v\n", testVar)
	testVar.ChangeSample("bb")
	prettyPrint(testVar)
	endTopic()

	newTopic("Loop")
	tasks()
	endTopic()

	newTopic("Interfaces")
	l := LoveMessagePrinter{}
	PrintSomething(l)
	endTopic()

	newTopic("go verb")
	goVerb()
	endTopic()

	newTopic("go verb - wait for all")
	waitForAll()
	endTopic()

	newTopic("channels")
	pong()
	endTopic()
}

/* output
 {Sample:aa SampleInt:1}
{
        "sample": "bb",
        "sample_int": 1
}
*/
