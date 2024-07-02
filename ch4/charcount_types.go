package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := map[string]int{
		"letters": 0,
		"numbers": 0,
		"spaces": 0,
		"punct": 0,
		"symbols": 0,
	}
	
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v/n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch{
		case unicode.IsLetter(r):
			counts["letters"]++
		case unicode.IsNumber(r):
			counts["numbers"]++
		case unicode.IsSpace(r):
			counts["spaces"]++
		case unicode.IsPunct(r):
			counts["punct"]++
		case unicode.IsSymbol(r):
			counts["symbols"]++
		}
	}
	fmt.Printf("type\tcount\n")
	for i, n := range counts {
		fmt.Printf("%s\t%d\n", i, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
