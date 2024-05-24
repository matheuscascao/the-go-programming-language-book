// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return buf.String()
	}
	for i, v := range s {
		buf.WriteRune(v)
		if i%3 == 0 {
			buf.WriteString(",")
		}
	}
	return buf.String()
	// return comma(s[:n-3]) + "," + s[n-3:]
}

//!-