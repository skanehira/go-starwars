package main

import (
	"bufio"
	"fmt"
	"strconv"
	"time"

	"flag"

	"github.com/gobuffalo/packr"
	colorable "github.com/mattn/go-colorable"
)

var (
	// animation speed option
	speed = flag.Int("s", 40, "speed, up speed by making it smaller")
	// use colorable to support windows
	stdOut = bufio.NewWriter(colorable.NewColorableStdout())
	// screen height
	height = 14
)

func clear() {
	// https://www.grapecity.com/developer/support/powernews/column/clang/047/page02.htm
	fmt.Fprint(stdOut, "\x1b[2J")
	fmt.Fprint(stdOut, "\x1b[1;1H")
	stdOut.Flush()
}

func main() {
	// using packr to embed text in binary
	file, err := packr.NewBox("./resources").Open("sw1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// parse option
	flag.Parse()

	// use default value
	var buffer string
	var duration int
	var i int

	// read a line from embed text file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if i%height == 0 {
			// get duration from text line
			duration, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}

			// print buffered line
			fmt.Println(buffer)

			time.Sleep(time.Duration((*speed)*duration) * time.Millisecond)

			// clear terminal screen, init buffer and duration
			clear()
			buffer = ""
			duration = 0
		} else {
			// text line into buffer
			buffer += fmt.Sprintln(scanner.Text())
		}
		i++
	}
}
