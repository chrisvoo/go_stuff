# Basics

The `go fmt` command will remove semicolons and adjust other formatting issues. The `go build` command compiles Go source code and produces an executable. To remove the output from the compilation process, run `go clean`.
The reason that most projects start with the `go mod init` command is that it simplifies the build process. Instead of specifying a particular code file, the project can be built and executed using a period, indicating the project in the current directory.
The `go vet` command identifies statements likely to be mistakes. Unlike a linter, which will often focus on style issues, the go vet command finds code that compiles but that probably won’t do what the developer intended.
The `go fmt` command formats Go source code files for consistency. There are no configuration options to change the formatting applied by the go fmt command, which transforms code into the style specified by the Go development team.

## Debugger

The standard debugger for Go applications is called Delve. By default, the dlv command will be installed in the `~/go/bin` folder (although this can be overridden by setting the GOPATH environment variable).

**Put ~/go/bin/ in your global path**

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

Procedure:

* `dlv debug <SCRIPT_NAME>` (eg: `main.go`)
* `break <BREAKPOINT_NAME> <PACKAGE.FUNCTION:LINE>` (eg `break bp1 main.main:3`)
* create a condition for the breakpoint so that execution will be halted only when a specified expression evaluates to true: `condition bp1 i == 2`
* `continue` starts/resume the execution
* `next`: moves to the next line
* `step`: steps into the current statement
* `stepout`: steps out of the current statement
* `restart` and then `continue`: restarts the process and begins the execution

The debugger provides a full set of commands for inspecting and altering the state of the application:

* `print <exp>`: valuates an expression and displays the result
* `set <variable> = <value>`: changes the value of the specified variable.
* `locals`: this command prints the value of all local variables
* `whatis <expr>`:  prints the type of the specified expression

These operations can be done by the plugin for Visual Studio code.

## Linting

```bash
go install github.com/mgechev/revive@latest
```

Example:

```text
~/GitHub/go_stuff/pro-go/basics:$ revive
main.go:1:1: should have a package comment
main.go:6:1: exported function PrintHello should have comment or be unexported
main.go:10:1: exported function PrintNumber should have comment or be unexported
```

The reason that the linter is so strict about comments is because they are used by the `go doc` command, which generates documentation from source code comments. If you don’t want to apply a rule at all, then you can use [a TOML-format configuration file](https://github.com/mgechev/revive?tab=readme-ov-file#configuration).
For VS Code it is easy to change the linter to revive using the Preferences ➤ Extensions ➤ Go ➤ Lint Tool configuration option. If you want to use a custom configuration file, use the Lint Flags configuration option to add a flag with the value -config=./revive.toml, which will select the revive.toml file.

## Packages and modules

* **define a package**: create a folder and add code files with `package` statements.
* **use a package**: add an import statement that specifies the path to the package
* **ACL to features in a package**: export features by using an initial uppercase letter in their names. Lowercase initial letters are unexpected and cannot be used outside the package
* **package conflicts**: use an alias or a dot import.
* **perform tasks when a package is loaded**: define an init function
* **execute a package init function without importing the features in contains**: use the blank identifier import statement.
* **use an external package**: use `go get` command
* **Remove unused package dependencies**: use `go mod tidy`

The name specified by the package statement should match the name of the folder in which the code files are created, which is store in this case.

The `go get` command adds dependencies to the go.mod file, but these are not removed automatically if the external package is no longer required. To update the go.mod file to reflect the change, run `go mod tidy`.

## Channels and goroutines

A **goroutine** is a lightweight thread. All Go programs use at least one goroutine because this is how Go executes the code in the main function. When compiled Go code is executed, the runtime creates a goroutine that starts executing the statements in the entry point, which is the main function in the main package. Each statement in the main function is executed in the order in which they are defined. The goroutine keeps executing statements until it reaches the end of the main function, at which point the application terminates.
The goroutine executes each statement in the main function synchronously, which means that it waits for the statement to complete before moving on to the next statement.
A goroutine allows to execute a function asynchronously. The result produced by such a function can be produced through a channel. Values can be sent and received using a channel through arrow expressions. The close function indicates that no
further values will be sent over the channel.
Go allows the developer to create additional goroutines, which execute code **at the same time** as the main goroutine.
Getting a result from a function that is being executed asynchronously can be complicated because it requires coordination between the goroutine that produces the result and the goroutine that consumes the result. To address this issue, Go provides **channels**, which are conduits through which data can be sent and received.

### Coordinating Channels

Receiving from a channel is a **blocking operation**, meaning that execution will not continue until a value has been received. By default, sending and receiving through a channel are blocking operations. This means a goroutine that sends a value will not execute any further statements until another goroutine receives the value from the channel. If a second goroutine sends a value, it will be blocked until the channel is cleared, causing a queue of goroutines waiting for values to be received. This happens in the other direction, too, so that goroutines that receive values will block until another goroutine sends one.
An alternative approach is to create **a channel with a buffer**, which is used to accept values from a sender and store them until a receiver becomes available. This makes sending a message a nonblocking operation, allowing a sender to pass its value to the channel and continue working without having to wait for a receiver.
The `select` keyword is used to group operations that will send or receive from channels, which allows for complex arrangements of goroutines and channels to be created. The simplest use for select statements is to **receive from a channel without blocking**, ensuring that a goroutine won’t have to wait when the channel is empty. Care must be taken to manage closed channels because they will provide a nil value for every receive operation that occurs after the channel has closed, relying on the closed indicator to show that the channel is closed. Unfortunately, this means that case statements for closed channels will always be chosen by select statements because they are always ready to provide a value without blocking, even though that value isn't useful. Managing closed channels requires two measures:
* to prevent the select statement from choosing a channel once it is closed by nullifying the channel variable. A nil channel is never ready and will not be chosen, allowing the select statement to move onto other case statements
* to break out of the for loop when all the channels are closed, without which the select statement would endlessly execute the default clause.

You can combine case statements with send and receive operations in the same select statement. When the select statement is executed, the Go runtime builds a combined list of case statements that can be executed without blocking and picks one at random, which can be either a send or a receive statement.

## Errors handling

If a function is being executed using a goroutine, then the only communication is through the channel, which means that details of any problems must be communicated alongside successful operations. An approach is to create a custom type that consolidates both outcomes.
In alternative, the `errors` package, which is part of the standard library, provides a `New` function that returns an error whose content is a string. The drawback of this approach is that it creates simple errors.

### panic

Some errors are so serious they should lead to the immediate termination of the application, a process known as panicking. When the panic function is called, the execution of the enclosing function is halted, and any defer functions are performed. The panic bubbles up through the call stack, terminating execution of the calling functions and invoking their defer functions. Go provides the built-in function `recover`, which can be called to stop a panic from working its way up the call stack and terminating the program.  
A panic works its way up the stack only to the top of the current goroutine, at which point it causes termination of the application. This restriction means that **panics must be recovered within the code that a goroutine executes**.