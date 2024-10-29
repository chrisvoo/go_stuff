package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	arg := "main"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	switch arg {
	case "main":
		{
			current := time.Now()
			// Local represents the system's local time zone (here system’s time zone)
			specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
			unix := time.Unix(1433228090, 0)
			PrintTime("Current", &current)
			PrintWithRefTime("Current with ref: ", &current)
			PrintTime("Specific", &specific)
			PrintWithRefTime("Specific with ref:", &specific)
			PrintTime("UNIX", &unix)
			PrintWithRefTime("UNIX with ref:", &unix)

			// Parse a time value from a string in the standard Unix format.
			t, err := time.Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
			if err != nil { // Always check errors even if they should not happen.
				panic(err)
			}

			// time.Time's Stringer method is useful without any format.
			fmt.Println("default format:", t)
			// Predefined constants in the package implement common layouts.
			fmt.Println("Unix format:", t.Format(time.UnixDate))
			// The time zone attached to the time value affects its output.
			fmt.Println("Same, in UTC:", t.UTC().Format(time.UnixDate))

			tz, err := time.LoadLocation("Asia/Shanghai")
			if err != nil { // Always check errors even if they should not happen.
				panic(err)
			}
			fmt.Println("in Shanghai with colon seconds:", t.In(tz).Format("2006-01-02T15:04:05 -07:00:00"))

			tim, err := time.Parse(time.RFC822, "09 Jun 95 04:59 BST")
			if err == nil {
				Printfln("After: %v", tim.After(time.Now()))
				Printfln("Round: %v", tim.Round(time.Hour))
				Printfln("Truncate: %v", tim.Truncate(time.Hour))
			} else {
				fmt.Println(err.Error())
			}

			t1, _ := time.Parse(time.RFC822Z, "09 Jun 95 04:59 +0100")
			t2, _ := time.Parse(time.RFC822Z, "08 Jun 95 23:59 -0400")
			Printfln("Equal Method: %v", t1.Equal(t2))

			// Creating and Inspecting a Duration
			var d time.Duration = time.Hour + (30 * time.Minute)
			Printfln("Hours: %v", d.Hours())
			Printfln("Mins: %v", d.Minutes())
			Printfln("Seconds: %v", d.Seconds())
			Printfln("Millseconds: %v", d.Milliseconds())
			rounded := d.Round(time.Hour)
			Printfln("Rounded Hours: %v", rounded.Hours())
			Printfln("Rounded Mins: %v", rounded.Minutes())
			trunc := d.Truncate(time.Hour)
			Printfln("Truncated  Hours: %v", trunc.Hours())
			Printfln("Rounded Mins: %v", trunc.Minutes())

			// Creating Durations Relative to Times
			toYears := func(d time.Duration) int {
				return int(d.Hours() / (24 * 365))
			}
			future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
			past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)
			Printfln("Future: %v", toYears(time.Until(future)))
			Printfln("Past: %v", toYears(time.Since(past)))

			dur, err := time.ParseDuration("1h30m")
			if err == nil {
				Printfln("Hours: %v", dur.Hours())
				Printfln("Mins: %v", dur.Minutes())
				Printfln("Seconds: %v", dur.Seconds())
				Printfln("Millseconds: %v", dur.Milliseconds())
			} else {
				fmt.Println(err.Error())
			}
		}
	case "after":
		{
			// The time package provides a small set of functions that are useful for working with goroutines and channels
			nameChannel := make(chan string)
			go WriteToChannel(nameChannel)
			time.AfterFunc(time.Second*5, func() {
				WriteToChannel(nameChannel)
			})
			for name := range nameChannel {
				Printfln("Read name: %v", name)
			}
		}
	case "notify":
		{
			nameChannel := make(chan string)
			go Notify(nameChannel)
			channelOpen := true
			for channelOpen {
				Printfln("Starting channel read")
				select {
				case name, ok := <-nameChannel:
					if !ok {
						channelOpen = false
						break
					} else {
						Printfln("Read name: %v", name)
					}
				case <-time.After(time.Second * 2):
					Printfln("Timeout")
				}
			}
		}
	case "timer":
		{
			nameChannel := make(chan string)
			go WriteWithTimer(nameChannel)
			for name := range nameChannel {
				Printfln("Read name: %v", name)
			}
		}
	case "recurring":
		{
			nameChannel := make(chan string)
			go RecurringNotify(nameChannel)
			for name := range nameChannel {
				Printfln("Read name: %v", name)
			}
		}
	case "ticker":
		{
			nameChannel := make(chan string)
			go WriteTicker(nameChannel)
			for name := range nameChannel {
				Printfln("Read name: %v", name)
			}
		}
	}
}
