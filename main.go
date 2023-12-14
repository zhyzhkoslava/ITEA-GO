package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func readFile(filePath string, linesCh chan []string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	linesCh <- lines
}

func main() {
	file1Path := "hw11/file1.txt"
	file2Path := "hw11/file2.txt"
	outputFilePath := "hw11/res.txt"

	var wg sync.WaitGroup
	linesCh := make(chan []string, 2)

	wg.Add(2)
	go readFile(file1Path, linesCh, &wg)
	go readFile(file2Path, linesCh, &wg)

	wg.Wait()
	close(linesCh)

	var combinedLines []string
	for lines := range linesCh {
		combinedLines = append(combinedLines, lines...)
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error opening output file:", err)
		return
	}
	defer outputFile.Close()

	for _, line := range combinedLines {
		_, err := fmt.Fprintln(outputFile, line)
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}
}
