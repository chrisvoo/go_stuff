# Getting started

When your code imports packages contained in other modules, you manage those 
dependencies through your code's own module. That module is defined by a go.mod 
file that tracks the modules that provide those packages. That go.mod file stays 
with your code, including in your source code repository.
Go code is grouped into packages, and packages are grouped into modules. Your 
module specifies dependencies needed to run your code, including the Go version 
and the set of other modules it requires.  

In Go, a function whose name starts with a capital letter can be called by a 
function not in the same package. This is known in Go as an exported name.

* `go mod tidy`: resolves and downloads new imported modules
* `go mod init <name>`: creates a go.mod file to track your code's dependencies.
* `go mod edit -replace greetings=../greeting`: point an unpublished module to a directory where it is placed
* `go test`: tests a module
* `go build`: builds the executable
* `go list -f '{{.Target}}'`: prints the GO Install path, where the go 
command will install the current package. As an alternative, if you already have
a directory like $HOME/bin in your shell path and you'd like to install your Go programs there,you can change the 
install target by setting the GOBIN variable using the go env command: `go env -w GOBIN=/path/to/your/bin`
* `go install`: install the package in the install path.