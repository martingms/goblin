// TODO: Support floating point numbers.
// TODO: Add support for -w.
// TODO: Print usage etc.

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	f = flag.String("f", "%v", "format string. See http://golang.org/pkg/fmt.")
	s = flag.String("s", "\n", "separator string.")
	//w = flag.Bool("w", false, "equalize width by padding with leading zeroes.")
)

func main() {
	flag.Parse()
	fmtStr := *f + *s

	var (
		first     = 1
		increment = 1
		last      int
		err       error
	)

	switch len(flag.Args()) {
	case 0:
		fmt.Fprintln(os.Stderr, "TODO: Print help.") //TODO
	case 1:
		last, err = strconv.Atoi(flag.Arg(0))
	case 2:
		first, err = strconv.Atoi(flag.Arg(0))
		last, err = strconv.Atoi(flag.Arg(1))
	case 3:
		first, err = strconv.Atoi(flag.Arg(0))
		increment, err = strconv.Atoi(flag.Arg(1))
		last, err = strconv.Atoi(flag.Arg(2))
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, string(os.Args[0])+": invalid argument.")
		fmt.Fprintln(os.Stderr, "Try '"+string(os.Args[0])+" --help' for more information.")
		os.Exit(1)
	}

	for ; first <= last; first += increment {
		fmt.Printf(fmtStr, first)
	}
}
