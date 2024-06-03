// repl/repl.go

package repl

import (
	"bufio"
	"fmt"
	"hippo/evaluator"
	"hippo/lexer"
	"hippo/object"
	"hippo/parser"
	"io"
)

const PROMPT = ">> "

const HIPPO = `⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⢤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⠖⠓⠲⢤⡀
⠀⠀⠀⠀⠀⠀⠀⣀⡾⡭⢤⡄⠈⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⡇⢠⡾⠛⢲⣿⡀
⠀⠀⠀⠀⠀⠀⠀⢹⣿⠁⢸⣿⠀⢻⡄⢀⣠⠤⠴⠶⠶⢦⡤⠤⠒⠒⠒⠒⠦⣤⡀⠀⠀⣸⠀⢸⡇⠀⢸⢿⡇
⠀⠀⠀⠀⠀⠀⠀⠸⣿⣆⠀⢿⣄⣤⠟⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡤⠴⠦⢭⣷⣶⠃⢀⡞⠀⣠⠋⡼
⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣷⢴⡿⠋⠉⠉⠓⠄⠀⠀⠀⠀⠀⠀⠀⡔⠁⠀⠀⠀⠀⠈⠻⣷⣿⠴⢊⡡⠞⠁
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⠁⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠀⠀⠤⣤⣤⡀⠀⠀⠀⢻⡚⠉
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⡿⠚⠛⠻⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⠟⠉⠉⠙⢦⠀⠀⠈⢧
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡼⢁⣴⣶⣤⢸⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⡟⢀⣾⣿⠷⣬⡇⠀⠀⠈⢳⡄
⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⡇⣼⣿⣧⣽⣾⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠈⣿⣷⣴⣿⡟⠀⠀⠀⠀⠹⡆
⠀⠀⠀⠀⠀⠀⠀⠀⡾⠁⣳⣙⣿⣿⣿⣃⡀⠀⠀⠀⠀⠀⣀⣀⡀⠀⠀⠳⣝⣿⡿⠟⠳⠀⠀⠀⠀⠀⢻
⠀⠀⠀⠀⠀⠀⠀⢸⡇⠀⣠⠿⠏⠉⠉⠉⠉⠉⠙⡿⠛⠉⠉⠀⠉⠁⠀⠀⠜⠁⠀⠀⠀⠀⠀⠀⠀⠀⠸⡇
⠀⠀⠀⠀⠀⠀⠀⠀⣧⡾⠁⢠⣿⠙⣆⠀⠀⠀⡼⠁⠀⠀⠀⠀⡴⢋⣿⣷⠀⠀⠀⢀⣶⡄⠀⠀⠀⠀⠀⣷
⠀⠀⠀⠀⠀⠀⠀⠀⣼⠁⠀⠈⠻⣿⡿⠀⠀⢠⡇⠀⠀⠀⠀⠐⢳⠿⠛⠉⠀⠀⠀⢸⣿⠃⠀⠀⠀⠀⢠⡟
⠀⠀⠀⠀⠀⠀⠀⠀⣿⠀⠀⠀⠀⠈⠁⠀⠀⢸⡇⠀⠀⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⠈⡇⠀⠀⠀⠀⢠⡾⠁
⠀⠀⠀⠀⠀⠀⠀⠀⠸⣆⠀⠀⠀⠀⠀⠋⠀⠘⣇⠀⠀⠀⠂⠀⠈⠀⠀⠀⠀⠀⠀⢸⠃⠀⠀⣠⡴⠋
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⣦⡀⠈⠁⠀⢀⣀⣶⣦⣤⣀⠀⠀⠉⠀⠐⠂⠀⠀⢀⣰⠏⢀⣤⣾⡉
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡿⣷⡒⠋⠉⠁⠀⣀⣀⠈⠙⠓⠶⠤⠤⠤⠴⠖⣋⣥⠞⠋
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀  ⠈⠙⠓⠒⠚⠋⠉⠉⠙⠲⢤⣄⣀⣀⣤⠴⠚⠯
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "exit." {
			break
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, HIPPO)
	io.WriteString(out, "Woops! We ran into some hippo business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
