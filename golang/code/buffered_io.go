package main

import (
	"bufio"
	"fmt"
	"os"
)

// START OMIT

// Wrap STDIN in buffered reader (default size: 4k).
var reader *bufio.Reader = bufio.NewReader(os.Stdin) // HL

// Wrap STDOUT in buffered writer (default size: 4k).
var writer *bufio.Writer = bufio.NewWriter(os.Stdout) // HL

// Helper function to Scan the buffered STDIN.
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) } // HL
// Helper function to Print to the buffered STDOUT.
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) } // HL

func main() {
	var a, b int
	scanf("%d %d\n", &a, &b) // HL
	printf("%d\n", a+b) // HL
	writer.Flush() // STDOUT must be flushed.
}

// END OMIT
