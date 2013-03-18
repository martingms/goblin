package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		printStdIn()
	}

	for _, f := range os.Args[1:] {
		printFile(f)
	}
}

func printStdIn() {
	reader := bufio.NewReader(os.Stdin)
	for {
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Print(s)
	}
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	out := bufio.NewWriter(os.Stdout)
	buf := make([]byte, 1024)

	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		_, err = out.Write(buf[:n])
		if err != nil {
			panic(err)
		}
	}

	err = out.Flush()
	if err != nil {
		panic(err)
	}
}
