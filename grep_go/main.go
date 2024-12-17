package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type GrepFlags struct {
	A, B, C       int
	c, i, v, F, n bool
}

var flags GrepFlags

func main() {
	searchStr, files := parseFlags()
	processFiles(searchStr, files)
}

func parseFlags() (string, []string) {
	flag.IntVar(&flags.A, "A", 0, "Print N lines after match")
	flag.IntVar(&flags.B, "B", 0, "Print N lines before match")
	flag.IntVar(&flags.C, "C", 0, "Print N lines around match (A+B)")
	flag.BoolVar(&flags.c, "c", false, "Print count of matching lines")
	flag.BoolVar(&flags.i, "i", false, "Ignore case in matching")
	flag.BoolVar(&flags.v, "v", false, "Invert match")
	flag.BoolVar(&flags.F, "F", false, "Exact string match")
	flag.BoolVar(&flags.n, "n", false, "Print line numbers")
	flag.Parse()

	if len(flag.Args()) < 2 {
		log.Fatal("Usage: program <pattern> <file1> [file2 ...]")
	}

	searchStr := flag.Arg(0)
	files := flag.Args()[1:]
	return searchStr, files
}

func processFiles(searchStr string, files []string) {
	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			log.Printf("Error opening file %s: %v", filename, err)
			continue
		}
		defer file.Close()

		lineCount := processFile(file, searchStr)
		if flags.c {
			printCount(lineCount, filename, len(files) > 1)
		}
	}
}

func processFile(file *os.File, searchStr string) int {
	scanner := bufio.NewScanner(file)
	var buffer []string
	lineCount, lineNumber := 0, 0
	afterCounter := 0

	if flags.C > 0 {
		flags.A, flags.B = flags.C, flags.C
	}

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		processedLine, processedStr := preprocessLine(line, searchStr)
		matched := matchLine(processedLine, processedStr)

		if matched {
			lineCount++
			printMatchedLine(buffer, line, lineNumber)
			afterCounter = flags.A
			buffer = nil
		} else {
			buffer = updateBeforeBuffer(buffer, line, lineNumber)
			printAfterLines(&afterCounter, line, lineNumber)
		}
	}
	return lineCount
}

func preprocessLine(line, searchStr string) (string, string) {
	if flags.i || flags.v {
		line = strings.ToLower(line)
		searchStr = strings.ToLower(searchStr)
	}
	return line, searchStr
}

func matchLine(line, searchStr string) bool {
	if flags.F {
		return line == searchStr
	}
	matched := strings.Contains(line, searchStr)
	if flags.v {
		return !matched
	}
	return matched
}

func printMatchedLine(buffer []string, line string, lineNumber int) {
	if flags.c {
		return
	}
	for _, bufLine := range buffer {
		fmt.Println(bufLine)
	}

	if flags.n {
		fmt.Printf("%d:%s\n", lineNumber, line)
	} else {
		fmt.Println(line)
	}
}

func updateBeforeBuffer(buffer []string, line string, lineNumber int) []string {
	if flags.B > 0 {
		if len(buffer) >= flags.B {
			buffer = buffer[1:]
		}
		if flags.n {
			buffer = append(buffer, fmt.Sprintf("%d:%s", lineNumber, line))
		} else {
			buffer = append(buffer, line)
		}
	}
	return buffer
}

func printAfterLines(afterCounter *int, line string, lineNumber int) {
	if *afterCounter > 0 {
		if flags.n {
			fmt.Printf("%d: %s\n", lineNumber, line)
		} else {
			fmt.Println(line)
		}
		*afterCounter--
	}
}

func printCount(lineCount int, filename string, multipleFiles bool) {
	if multipleFiles {
		fmt.Printf("%s:%d\n", filename, lineCount)
	} else {
		fmt.Printf("%d\n", lineCount)
	}
}
