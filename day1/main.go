package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	sum := 0

	for fileScanner.Scan() {
		first := -1
		last := -1
		cstr := ""

		line := fileScanner.Text()
		for _, ch := range line {
			cstr = string(ch)
			if v, err := strconv.Atoi(cstr); err == nil {
				if first == -1 {
					first = v
				} else {
					last = v
				}
			}

		}
		if last == -1 {
			last = first
		}
		if first != -1 {
			sum += first*10 + last
		}
	}

	fmt.Printf("sum %v", sum)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
