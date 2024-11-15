package main

import (
	"fmt"
	"io"
	"strings"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func ScanFromReader(
	reader io.Reader,
	template string,
	vals ...interface{}) (int, error) {
	return fmt.Fscanf(reader, template, vals...)
}

func ScanSingle(reader io.Reader, val interface{}) (int, error) {
	return fmt.Fscan(reader, val)
}

func WriteFormatted(writer io.Writer, template string, vals ...interface{}) {
	fmt.Fprintf(writer, template, vals...)
}

func WriteReplaced(writer io.Writer, str string, subs ...string) {
    replacer := strings.NewReplacer(subs...)
    replacer.WriteString(writer, str)
}

func ProcessData(reader io.Reader) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		// used to signal when the Reader reaches the end of the data
		if err == io.EOF {
			break
		}
	}
}

func WriteData(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			writer.Write(b[0:count])
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func CopyData(reader io.Reader, writer io.Writer) {
	/*
	   func io.Copy(dst io.Writer, src io.Reader) (written int64, err error)
	   Copy copies from src to dst until either EOF is reached on src or an error occurs.
	   It returns the number of bytes copied and the first error encountered while copying,
	   if any.
	*/
	count, err := io.Copy(writer, reader)
	if err == nil {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

/* Defines a Writer parameter, which it uses to write bytes from a string. */
func GenerateData(writer io.Writer) {
	data := []byte("Kayak, Lifejacket")
	writeSize := 4
	for i := 0; i < len(data); i += writeSize {
		end := i + writeSize
		if end > len(data) {
			end = len(data)
		}
		count, err := writer.Write(data[i:end])
		Printfln("Wrote %v byte(s): %v", count, string(data[i:end]))
		if err != nil {
			Printfln("Error: %v", err.Error())
		}
	}
	if closer, ok := writer.(io.Closer); ok {
		closer.Close()
	}
}

/* Defines a Reader parameter, which it uses to read bytes of data, which are then used to create a string */
func ConsumeData(reader io.Reader) {
	data := make([]byte, 0, 10)
	slice := make([]byte, 2)
	for {
		count, err := reader.Read(slice)
		if count > 0 {
			Printfln("Read data: %v", string(slice[0:count]))
			data = append(data, slice[0:count]...)
		}
		if err == io.EOF {
			break
		}
	}
	Printfln("Read data: %v", string(data))
}
