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
	speed  = flag.Int("s", 40, "speed, up speed by making it smaller")
	stdOut = bufio.NewWriter(colorable.NewColorableStdout())
)

func clear() {
	fmt.Fprint(stdOut, "\x1b[2J")
	fmt.Fprint(stdOut, "\x1b[0;0H")
	stdOut.Flush()
}

func parse() {
	flag.Parse()
}

func main() {
	file, err := packr.NewBox("./resources").Open("sw1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	parse()

	var buffer []string
	var duration int

	height := 14
	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		if i%height == 0 {
			duration, err = strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}

			for _, buf := range buffer {
				fmt.Println(buf)
			}

			time.Sleep(time.Duration((*speed)*duration) * time.Millisecond)
			clear()
			buffer = []string{}
			duration = 0
		} else {
			buffer = append(buffer, scanner.Text())
		}
		i++
	}
}
