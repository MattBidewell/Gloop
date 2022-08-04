# Gloop programming language

> Do not use in production. Like seriously. I mean you can... but I wouldnt. At least not yet.

Gloop is a experiment in learning how interpreters, compilers work and ultimately languages work.

This project should be used as a case study and not a implementable language. Follow for a journey in creating an interpreter. Updates will be provided below.

### Testing
Run `go test /{folder}` to test the individual folder

You can also run `sh ./test_runner.sh` which will explore all folders and run the tests files if they exist.

Heavily inspired by [Thorsten Ball's Writing an Interpreter in Go](https://interpreterbook.com/)

### Updates

#### 0.0.4
- Added parsing of let and return statements.
- The rest of the language wil consist of expressions.
#### 0.0.3
`01/07/2022`
- Added a basic parser as a step forward into making this a programming language.
  - Only parses let statments for the time being.
  - Is a recursive descent parser

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