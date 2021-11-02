package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	var filename = "tools.go"

	log := func(v ...interface{}) {
		_, _ = fmt.Fprintln(os.Stderr, v...)
	}

	if len(os.Args) == 2 {
		if os.Args[1] == "-h" {
			log("usage:   go-tool-installer <path-to-tools-file")
			log("example: go-tool-installer tools/tools.go")
			os.Exit(1)
		}
		filename = os.Args[1]
	}

	fset := token.NewFileSet()

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, filename, nil, parser.ImportsOnly)
	if err != nil {
		// don't care about the possible error/bytes written;
		// we're already bailing from the application.
		log(err)
		os.Exit(1)
	}

	fmt.Print("go install")
	for _, s := range f.Imports {
		fmt.Print(" " + s.Path.Value)
	}
	fmt.Println()
}
