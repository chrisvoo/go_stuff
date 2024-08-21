package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Custom type. Structs allow a set of related values to be grouped together
type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

/*
make allocates an underlying array of size 10 and returns a slice of length 0 and
capacity 10 that is backed by this underlying array.
A slice is a variable-length array. Slices are resized automatically as new items
are added, and the initial capacity determines how many items can be added before
the slice has to be resized. In this case, ten items can be added to the slice
before it has to be resized.
[] denotes a slice, * denotes a pointer. By specifying that my slice will contain pointers,
I am telling Go not to create copies of my Rsvp values when I add them to the slice.
*/
var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3) // a map whose keys are strings, and values pointers to Template

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	// The range keyword is used with the for keyword to enumerate arrays, slices, and maps.
	for index, name := range templateNames {
		t, err := template.ParseFiles("templates/layout.html", "templates/"+name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			// can be called when an unrecoverable error happens
			panic(err)
		}
	}
}

func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}

type formData struct {
	*Rsvp
	Errors []string
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templates["form"].Execute(writer, formData{
			// Rsvp empty, Go doesn’t have a new keyword, and values are created using braces
			// default values could be set here. The ampersand (the & character) creates a pointer to a value
			Rsvp: &Rsvp{}, Errors: []string{},
		})
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		responseData := Rsvp{
			Name:       request.Form["name"][0],
			Email:      request.Form["email"][0],
			Phone:      request.Form["phone"][0],
			WillAttend: request.Form["willattend"][0] == "true",
		}
		errors := []string{}
		if responseData.Name == "" {
			errors = append(errors, "Please enter your name")
		}
		if responseData.Email == "" {
			errors = append(errors, "Please enter your email address")
		}
		if responseData.Phone == "" {
			errors = append(errors, "Please enter your phone number")
		}
		if len(errors) > 0 {
			templates["form"].Execute(writer, formData{
				Rsvp: &responseData, Errors: errors,
			})
		} else {
			// The append function is used to append a value to a slice.  If I had not used a pointer, then my
			// Rsvp value would be duplicated when it is added to the slice.
			responses = append(responses, &responseData)
			if responseData.WillAttend {
				templates["thanks"].Execute(writer, responseData.Name)
			} else {
				templates["sorry"].Execute(writer, responseData.Name)
			}
		}
	}
}

// entry point for the application
func main() {
	loadTemplates()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":5500", nil)
	if err != nil {
		fmt.Println(err)
	}
}
