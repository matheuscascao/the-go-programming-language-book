package main

import (
	"fmt"
	"os"
	"strconv"

	"./conv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := conv.Fahrenheit(t) // Using conv package's Fahrenheit type
		c := conv.Celsius(t)    // Using conv package's Celsius type
		fmt.Printf("%s = %s, %s = %s\n",
			f, conv.FToC(f), c, conv.CToF(c)) // Accessing functions from conv package
	}
}