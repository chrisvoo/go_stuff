# Each possible operation is called a target. You can invoke a target by 
# running make <target> in the terminal.
# See https://makefiletutorial.com/

# Defines which target is run when no target is specified.
.DEFAULT_GOAL := build

# The .PHONY target is a special target that tells make that the listed 
# targets are not files.
.PHONY:fmt vet build clean
fmt:
	go fmt ./...

# vet. must run fmt target first, so it is listed as a dependency.
vet: fmt
	go vet ./...

build: vet
	go build

clean:
	go clean