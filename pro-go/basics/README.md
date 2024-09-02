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

