package endpoint

import (
	"encoding/json"
	"interpreter/evaluator"
	"interpreter/lexer"
	"interpreter/object"
	"interpreter/parser"
	"io"
	"net/http"
)

type RequestData struct {
	Code string `json:"code"`
}

func RunCodeHandler(w http.ResponseWriter, r *http.Request) {
	env := object.NewEnvironment()

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	l := lexer.New(data.Code)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		printParserErrors(w, p.Errors())
		return 
	}

	evaluated := evaluator.Eval(program, env)

	if evaluated != nil {
		w.WriteHeader(http.StatusOK)
    
       w.Write([]byte(evaluated.Inspect()))
	} else {
		http.Error(w, "Execution resulted in null output", http.StatusInternalServerError)
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
