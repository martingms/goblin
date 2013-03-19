package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	c = flag.Int("c", 0, "output the last N bytes") // TODO: "Alternatively, use -c +X to output everything FROM X"
	f = flag.Bool("f", false, "output appended data as the file grows")
	n = flag.Int("n", 10, "output the last N lines") // TODO: "Alternatively blah blah.
	s = flag.String("s", "1s", "with -f, sleep for duration between each call. See http://golang.org/pkg/time for format")

	client = new(http.Client)

	buf           bytes.Buffer
	lastModified  string
	currentOffset = "bytes=0-"
)

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "TODO: Print help.") //TODO
	}

	dur, err := time.ParseDuration(*s)
	if err != nil {
		panic(err)
	}

	url := flag.Arg(0)
	handleUrl(url, printTail)

	if *f {
		for {
			time.Sleep(dur)
			handleUrl(url, func(buf bytes.Buffer) {
				fmt.Print(buf.String())
			})
		}
	}
}

func handleUrl(url string, printFunc func(bytes.Buffer)) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("If-Modified-Since", lastModified)
	req.Header.Add("Range", currentOffset)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// 304 = Not modified.
	if resp.StatusCode == 304 {
		return
	}

	lastModified = resp.Header["Last-Modified"][0]
	currentOffset = "bytes=" + resp.Header["Content-Length"][0] + "-"

	defer resp.Body.Close()

	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		panic(err)
	}

	printFunc(buf)
	buf.Reset() // Reuse the buffer instead of making new ones.
}

func printTail(buf bytes.Buffer) {
	// If -c flag is not set (we're dealing with lines).
	if *c == 0 {
		lines := strings.Split(buf.String(), "\n")

		// FIXME: +-1 part kind of hacky. Right way to do it?
		for _, val := range lines[len(lines)-(*n+1) : len(lines)-1] {
			fmt.Println(val)
		}
	} else {
		bytes := buf.Bytes()

		fmt.Print(string(bytes[len(bytes)-*c : len(bytes)]))
	}
}
