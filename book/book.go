package book

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ScannerTest() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}

func MaxMin() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one more arguments")
		os.Exit(1)
	}
	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}

func TestError() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one more argument!"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, "This is Standart output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
