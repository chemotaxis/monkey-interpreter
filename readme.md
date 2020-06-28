# Monkey Interpreter

This directory contains an implementation of an interpreter for the Monkey
language as described in "Writing an Interpreter in Go" by Thorsten Ball.  The
main objective is to learn how an interpreter works.

## Running the Interpreter

- If available, click or run the binary.  REPL should appear in a terminal.

Otherwise, if you want to run using source:

- Install Go
- Download this directory
- Then, do this:

```shell
cd monkey
go run main.go
```

To run tests:

```shell
# Run all tests
cd monkey
go test ./...

# Run specific tests
go test ./<dir name or file name>
```

## Development

For the most part, I'm typing everything in the book as I'm reading (not copying
and pasting).  I'm trying to avoid looking at the reference code that comes with
the book in order to see what kinds of errors come up.

For each section of the book, I make a new branch.  Once everything is
implemented in that section, I merge without fast-forwarding and create a tag on
the merge commit, naming it after the section.

Roughly, development starts with adding tests, implementing the necessary code,
and debugging until the tests pass.

Although most of the code is right from the book, I do refactor and write
additional code.  I also plan to add other features.
