package main

import (
	"bytes"
	"flag"
	"fmt"
)

var n = flag.Bool("n", false, "suppresses the newline")

func main() {
	flag.Parse()

	var buf bytes.Buffer
	for i, str := range flag.Args() {
		buf.WriteString(str)
		if i < flag.NArg()-1 {
			buf.WriteString(" ")
		}
	}

	if !*n {
		buf.WriteString("\n")
	}

	_, err := fmt.Print(buf.String())

	if err != nil {
		panic(err)
	}
}
