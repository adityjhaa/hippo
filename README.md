# HIPPO

**An interpreter for the Hippo programming language built in *Go*!**
<hr>

<div align = "justify">
I'm currently reading <a href="https://interpreterbook.com/"><b>Writing an Interpreter in Go</b></a> by <b>Thorsten Ball</b>. Ball believes "<i>Interpreters are Magical!</i>" and so are they. The book covers fundamental concepts of interpreters, such as lexical analysis, parsing, and evaluation. We'll follow Ball's approach and work on each part in stages.


#### Stage 1 : The Lexer

We'll first create a lexer that takes as input our source code, does the lexical analysis and returns the tokens it creates along the way. This makes it easier for us to work with. It'll turn our <b>Hippo</b> code into the tokens designed by us.

The finished Stage 1 can be found <a href="https://github.com/adityjhaa/hippo_interpreter/tree/538fffca7567e8c4deb1eb50789e9b4c644b02e5">here</a>.

#### Stage 2 : The Parser

After creating the tokens, we have the parser, that uses these tokens to build up the structure of the program (commonly known as the abstract syntax tree). This AST also throws errors (if any) related to syntax and structure of <b>Hippo</b> code.

The finished Stage 2 can be found <a href="https://github.com/adityjhaa/hippo_interpreter/tree/0d46b5964c2d199928903d3332c867af617c1270">here</a>.

</div>

<hr>

#### Usage
```bash
cd hippo_interpreter
go mod init hippo
go mod tidy
go run main.go
```

<hr>
