package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	arg := "enc"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	switch arg {
	case "enc":
		var b bool = true
		var str string = "Hello"
		var fval float64 = 99.99
		var ival int = 200
		var pointer *int = &ival
		var writer strings.Builder
		encoder := json.NewEncoder(&writer)
		for _, val := range []interface{}{b, str, fval, ival, pointer} {
			encoder.Encode(val)
		}
		fmt.Print(writer.String())

		names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
		numbers := [3]int{10, 20, 30}
		var byteArray [5]byte
		copy(byteArray[0:], []byte(names[0]))
		byteSlice := []byte(names[0])

		writer.Reset()
		encoder.Encode(names)
		encoder.Encode(numbers)
		encoder.Encode(byteArray)
		encoder.Encode(byteSlice) // byte slices are expressed as base64-encoded strings
		fmt.Print(writer.String())

		m := map[string]float64{
			"Kayak":      279,
			"Lifejacket": 49.95,
		}
		writer.Reset()
		encoder.Encode(m)
		fmt.Print(writer.String())

		// Unexported fields are ignored
		Kayak := &Product{Name: "Kayak", Category: "Watersports", price: 589}
		writer.Reset()
		encoder.Encode(Kayak)
		fmt.Print(writer.String())

		writer.Reset()
		dp := DiscountedProduct{
			Product:  Kayak,
			Discount: 10.50,
		}
		/*
			encodes a pointer to the struct value. The Encode function follows the pointer and encodes the value at its location.
			We have also used a custom serialization implementing MarshalJSON
		*/
		encoder.Encode(&dp)
		fmt.Print(writer.String())

		writer.Reset()
		cdp := CustomDiscountedProduct{
			Product:  Kayak,
			Discount: 10.50,
			color:    "red",
		}
		encoder.Encode(&cdp)
		fmt.Print(writer.String())
	case "dec":
		// the JSON specification allows values to be separated by spaces or newline characters
		reader := strings.NewReader(`true "Hello" 99.99 200 [10,20,30] ["Kayak","Lifejacket",279]`)
		/* The Decoder is able to select the appropriate Go data type for JSON values, and this is
		achieved by providing a pointer to an empty interface as the argument to the Decode method */
		vals := []interface{}{}
		decoder := json.NewDecoder(reader)
		/* JSON uses a single data type to represent both floating-point and integer values. The Decoder decodes these numeric values as float64
		This behavior can be changed by calling the UseNumber method on the Decoder, which causes JSON number values to be decoded into the Number type.
		*/
		decoder.UseNumber()
		for {
			var decodedVal interface{} // will contain the output
			err := decoder.Decode(&decodedVal)
			if err != nil {
				if err != io.EOF {
					Printfln("Error: %v", err.Error())
				}
				break
			}
			vals = append(vals, decodedVal)
		}
		for _, val := range vals {
			/*  If attempting to convert to an integer fails, then the Float64 method can be called. If a number cannot be converted to either Go type,
			then the String method can be used to get the unconverted string from the JSON data */
			if num, ok := val.(json.Number); ok {
				if ival, err := num.Int64(); err == nil {
					Printfln("Decoded Integer: %v", ival)
				} else if fpval, err := num.Float64(); err == nil {
					Printfln("Decoded Floating Point: %v", fpval)
				} else {
					Printfln("Decoded String: %v", num.String())
				}
			} else {
				Printfln("Decoded (%T): %v", val, val)
			}
		}

		/* If you know the structure of the JSON data you are decoding, you can direct the Decoder to use specific Go types by using variables
		of that type to receive a decoded value */
		reader = strings.NewReader(`true "Hello" 99.99 200 [10,20,30] ["Kayak","Lifejacket",279] {"Kayak" : 279, "Lifejacket" : 49.95}`)
		var bval bool
		var sval string
		var fpval float64
		var ival int
		ints := []int{}
		mixed := []interface{}{} // generic
		amap := map[string]float64{}
		vals = []interface{}{&bval, &sval, &fpval, &ival, &ints, &mixed, &amap}
		decoder = json.NewDecoder(reader)
		for i := 0; i < len(vals); i++ {
			err := decoder.Decode(vals[i])
			if err != nil {
				Printfln("Error: %v", err.Error())
				break
			}
		}
		Printfln("Decoded (%T): %v", bval, bval)
		Printfln("Decoded (%T): %v", sval, sval)
		Printfln("Decoded (%T): %v", fpval, fpval)
		Printfln("Decoded (%T): %v", ival, ival)
		Printfln("Decoded (%T): %v", ints, ints)
		Printfln("Decoded (%T): %v", mixed, mixed)
		Printfln("Decoded (%T): %v", amap, amap)
	case "struct":
		// case-insensitive match, will ignore any struct field for which there is no JSON key
		reader := strings.NewReader(`
			{"Name":"Kayak","Category":"Watersports","Price":279}
			{"Name":"Lifejacket","Category":"Watersports" }
			{"name":"Canoe","category":"Watersports", "price": 100, "inStock": true }
		`)
		decoder := json.NewDecoder(reader)
		// decoder.DisallowUnknownFields() would produce Error: json: unknown field "inStock"
		for {
			var val Product
			err := decoder.Decode(&val)
			if err != nil {
				if err != io.EOF {
					Printfln("Error: %v", err.Error())
				}
				break
			} else {
				Printfln("Name: %v, Category: %v, Price: %v",
					val.Name, val.Category, val.price)
			}
		}
	}
}
