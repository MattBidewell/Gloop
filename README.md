# gloop-terpreter

## Interpreter
Gloop is a experiment in learning how interpreters and compilers work.

This project should be used as a case study and not a implementable language.

Follow for a journey in creating an interpreter. Updates will be provided below.

### Testing
Run `go test ./lexer` to test the lexical analysis interpreter.

Heavily inspired by [Thorsten Ball's Writing an Interpreter in Go](https://interpreterbook.com/)

### Updates

#### 0.0.2
`10/04/2022`
- Added more complex tokens, including support for
  - `==`
  - `!=`
  - `>=`
  - `<=`
- Added a repl for interaction with the lexical analysis interpreter.
#### 0.0.1
`06/04/2022`
- Initalised the project
- Created a simple lexical analysis interpreter to be used with simple tokens.