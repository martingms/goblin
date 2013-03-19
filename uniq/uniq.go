package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// TODO include support for skipping characters and fields

var u = flag.Bool("u", false, "Print unique lines.")
var d = flag.Bool("d", false, "Print (one copy of) duplicated lines.")
var c = flag.Bool("c", false, "Prefix a repetition count and a tab to each output line.  Implies -u and -d.")

func main() {
	flag.Parse()

	var reader *bufio.Reader
	if flag.NArg() == 0 {
		reader = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	}

	var lastline string
	occurrences := -1

	for {
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if s == lastline {
			occurrences++
		} else {
			if occurrences != -1 {
				printUniqLine(lastline, occurrences)
			}
			lastline = s
			occurrences = 1
		}
	}
	printUniqLine(lastline, occurrences)
}

func printUniqLine(line string, occurrences int) {
	if *c {
		line = fmt.Sprintf("\t%d %s", occurrences, line)
	}

	if *u && !*d {
		if occurrences == 1 {
			fmt.Print(line)
		}
	} else if *d && !*u {
		if occurrences > 1 {
			fmt.Print(line)
		}
	} else {
		fmt.Print(line)
	}
}
