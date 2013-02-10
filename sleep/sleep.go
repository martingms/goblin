package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: sleep DURATION")
		fmt.Fprintln(os.Stderr, "Sleep for DURATION seconds.")
		fmt.Fprintln(os.Stderr, "Optionally use suffixes, like 1s or 2h45m30s etcetera.")
		fmt.Fprintln(os.Stderr, "Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.")
		os.Exit(1)
	}

	dur, err := time.ParseDuration(os.Args[1])
	if err != nil {
		// Most likely missing unit.
		dur, err = time.ParseDuration(os.Args[1] + "s")
		if err != nil {
			panic(err)
		}
	}

	time.Sleep(dur)
}
