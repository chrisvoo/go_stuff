package main

import (
	"fmt"
	"time"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func PrintTime(label string, t *time.Time) {
	Printfln("%s: Day: %v: Month: %v Year: %v", label, t.Day(), t.Month(), t.Year())
}

func PrintWithRefTime(label string, t *time.Time) {
	// constant -> const time.UnixDate untyped string = "Mon Jan _2 15:04:05 MST 2006"
	layout := "2006-01-02 15:04:05"
	fmt.Println(label, t.Format(layout))
}

func WriteToChannel(channel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		//time.Sleep(time.Second * 1)
	}
	close(channel)
}

func Notify(channel chan<- string) {
	Printfln("Waiting for initial duration...")
	<-time.After(time.Second * 2)
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		time.Sleep(time.Second * 3)
	}
	close(channel)
}

func WriteWithTimer(channel chan<- string) {
	// NewTimer create a Timer that is reset before the specified duration elapses
	timer := time.NewTimer(time.Minute * 10)
	go func() {
		time.Sleep(time.Second * 2)
		Printfln("Resetting timer")
		timer.Reset(time.Second) // resets the timer so its duration is two seconds
	}()
	Printfln("Waiting for initial duration...")
	<-timer.C
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		//time.Sleep(time.Second * 3)
	}
	close(channel)
}

/*
As before, the utility of the channel created by the Tick function isn’t the Time
values sent over it, but the periodicity at which they are sent. The channel blocks
when there is no value to read, which allows channels created with the Tick function
to control the rate at which the RecurringNotify function generates values.
*/
func RecurringNotify(nameChannel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	tickChannel := time.Tick(time.Second)
	index := 0
	for {
		<-tickChannel
		nameChannel <- names[index]
		index++
		if index == len(names) {
			index = 0
		}
	}
}

func WriteTicker(nameChannel chan <- string) {
    names := []string { "Alice", "Bob", "Charlie", "Dora" }
    ticker := time.NewTicker(time.Second / 10)
    index := 0
    for {
        <- ticker.C
        nameChannel <- names[index]
        index++
        if (index == len(names)) {
            ticker.Stop()
            close(nameChannel)
            break
        }
    }
}
