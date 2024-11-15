package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	arg := "main"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	switch arg {
	case "main":
		{
			// create a reader based on a string
			r := strings.NewReader("Kayak")
			ProcessData(r)

			r.Reset("Kayak")
			var builder strings.Builder
			WriteData(r, &builder)
			Printfln("String builder contents: %s", builder.String())

			r.Reset("Kayak")
			builder.Reset()
			CopyData(r, &builder)
			Printfln("String builder contents: %s", builder.String())
		}
	case "pipe":
		{
			// Pipes are used to connect code that consumes data through a Reader and code
			// that produces code through a Writer. Pipes are synchronous, such that the
			// PipeWriter.Write method will block until the data is read from the pipe.
			// This means that the PipeWriter needs to be used in a different goroutine from
			// the reader to prevent the application from deadlocking
			pipeReader, pipeWriter := io.Pipe()
			go GenerateData(pipeWriter)
			ConsumeData(pipeReader)
		}
	case "multi":
		{
			/* The Reader returned by the MultiReader function responds to the Read method
			with content from one of the underlying Reader values. When the first Reader returns
			EOF, then content is read from the second Reader. This process continues until the
			final underlying Reader returns EOF. */
			r1 := strings.NewReader("Kayak")
			r2 := strings.NewReader("Lifejacket")
			r3 := strings.NewReader("Canoe")
			concatReader := io.MultiReader(r1, r2, r3)
			ConsumeData(concatReader)

			var w1, w2, w3 strings.Builder
			combinedWriter := io.MultiWriter(&w1, &w2, &w3)
			GenerateData(combinedWriter)
			Printfln("Writer #1: %v", w1.String())
			Printfln("Writer #2: %v", w2.String())
			Printfln("Writer #3: %v", w3.String())
		}
	case "echo":
		{
			// The TeeReader function is used to create a Reader that will echo data to a strings.Builder
			r1 := strings.NewReader("Kayak")
			r2 := strings.NewReader("Lifejacket")
			r3 := strings.NewReader("Canoe")
			concatReader := io.MultiReader(r1, r2, r3)
			var writer strings.Builder
			teeReader := io.TeeReader(concatReader, &writer)
			ConsumeData(teeReader)
			Printfln("Echo data: %v", writer.String())
		}
	case "buf":
		{
			/* There are two reads that did not obtain 5 bytes of data. The penultimate read
			produced three bytes because the source data isn’t neatly divisible by five and there
			were three bytes of data left over. The final read returned zero bytes but received
			 the EOF error, indicating that the end of the data had been reached. */
			text := "It was a boat. A small boat."
			var customReader io.Reader = NewCustomReader(strings.NewReader(text))
			var writer strings.Builder
			slice := make([]byte, 5)
			// The default buffer size is 4,096 bytes
			buffered := bufio.NewReader(customReader)
			for {
				count, err := buffered.Read(slice)
				if count > 0 {
					Printfln("Buffer size: %v, buffered: %v", buffered.Size(), buffered.Buffered())
					writer.Write(slice[0:count])
				}
				if err != nil {
					break
				}
			}
			Printfln("Read data: %v", writer.String())

			var builder strings.Builder
			var bufWriter = bufio.NewWriterSize(NewCustomWriter(&builder), 20)
			for i := 0; true; {
				end := i + 5
				if end >= len(text) {
					bufWriter.Write([]byte(text[i:]))
					bufWriter.Flush()
					break
				}
				bufWriter.Write([]byte(text[i:end]))
				i = end
			}
			Printfln("Written data: %v", builder.String())
		}
	case "scan":
		{
			// Scanning Values from a Reader
			text := "Kayak Watersports $279.00"
			reader := strings.NewReader(text)
			var name, category string
			var price float64
			scanTemplate := "%s %s $%f"
			_, err := ScanFromReader(reader, scanTemplate, &name, &category, &price)
			if err != nil {
				Printfln("Error: %v", err.Error())
			} else {
				Printfln("Name: %v", name)
				Printfln("Category: %v", category)
				Printfln("Price: %.2f", price)
			}

			reader.Reset(text)
			for {
				var str string
				_, err := ScanSingle(reader, &str)
				if err != nil {
					if err != io.EOF {
						Printfln("Error: %v", err.Error())
					}
					break
				}
				Printfln("Value: %v", str)
			}

			// Writing Formatted Strings to a Writer
			var writer strings.Builder
			template := "Name: %s, Category: %s, Price: $%.2f"
			WriteFormatted(&writer, template, "Kayak", "Watersports", float64(279))
			fmt.Println(writer.String())

			// Using a Replacer with a Writer
			writer.Reset()
			text = "It was a boat. A small boat."
			subs := []string{"boat", "kayak", "small", "huge"}
			WriteReplaced(&writer, text, subs...)
			fmt.Println(writer.String())
		}
	}
}
