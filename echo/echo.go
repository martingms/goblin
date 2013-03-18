package main

import (
	"bytes"
	"flag"
	"fmt"
)

var n = flag.Bool("n", false, "suppresses the newline")
var e = flag.Bool("e", false, "enable interpretation of backslash escapes")

func main() {
	flag.Parse()

	var buf bytes.Buffer
	for i, str := range flag.Args() {
		buf.WriteString(str)
		if i < flag.NArg()-1 {
			buf.WriteString(" ")
		}
	}

	str := buf.String()

	if *e {
		for i := 0; i < len(str); i++ {
			c := str[i : i+1]
			if c == "\\" { // Forcing one-rune string instead of rune value.
				switch str[i+1] {
				case 'a':
					c = "\a"
				case 'b':
					c = "\b"
				case 'c':
					break
				case 'e':
					c = "\x1B"
				case 'f':
					c = "\f"
				case 'n':
					c = "\n"
				case 'r':
					c = "\r"
				case 't':
					c = "\t"
				case 'v':
					c = "\v"
				//TODO Output byte with octal value
				case '0':
				//TODO Output byte with hex value
				case 'x':
				}
				i++
			}
			fmt.Print(c)
		}
	} else {
		fmt.Print(str)
	}

	if !*n {
		fmt.Print("\n")
	}
}
