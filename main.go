package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func printUsage() {
	fmt.Printf("Usage: %v [SUBSTRING]\n", os.Args[0])
	fmt.Printf("  %v is intended to work with pipes and searches SUBSTRINGs in consecutive lines.\n", os.Args[0])
	fmt.Printf("  The last SUBSTRING is searched in all consecutive lines that are printed when it is found..\n")
	fmt.Printf("Example: 'ls -al | %v . ..'\n", os.Args[0])
	fmt.Printf("  searches for lines containing \".\" and multiple lines containing \".\" in first and \"..\" in all consecutive lines.\n")
}

func main() {
	fi, _ := os.Stdin.Stat() // get the FileInfo struct describing the standard input.

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		argLen := len(os.Args)
		if argLen < 2 {
			printUsage()
			return
		}

		reader := bufio.NewReader(os.Stdin)
		var compareString string
		var compareArgIndex int
		matchLine := 0

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if errors.Is(err, io.EOF) {
					line = line + "\n"
				}
			}
			if matchLine <= (argLen - 2) {
				compareArgIndex = 1 + matchLine
			} else {
				compareArgIndex = argLen - 1
			}
			compareString = os.Args[compareArgIndex]
			if strings.Contains(line, compareString) {
				print(line)
				matchLine++
			} else {
				matchLine = 0
			}

			if err != nil {
				break
			}
		}

	} else {
		printUsage()
	}
}
