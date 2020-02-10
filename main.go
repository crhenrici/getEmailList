package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var fName string
	fmt.Println("Name of file")
	fmt.Scanf("%s\n", &fName)
	fmt.Printf("Opening file %s\n", fName)
	dat, err := os.Open(fName)
	check(err)
	columns := readFile(dat)
	fmt.Println("entering names method")
	names := getNames(columns)
	dat.Close()
	fmt.Println("creating new file")
	newFile, err := os.Create("EmailList.txt")
	check(err)
	fmt.Println("entering writeFile method")
	writeFile(names, newFile)
	newFile.Close()
	fmt.Println("New File Created!")

}

func check(e error) {
	if e != nil {
		log.Fatalf("Failed openning file: %s", e)
	}
}

func readFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textLines []string
	var columns []string
	i := 0
	for scanner.Scan() {
		textLines = append(textLines, scanner.Text())
	}
	for i = 0; i < len(textLines); i++ {
		columns = append(strings.Split(textLines[i%6], "	"))
	}
	fmt.Println(columns)
	return columns
}

func getNames(columns []string) []string {
	length := len(columns)
	var names []string

	for i := 0; i < length; i++ {
		if (i % 7) == 0 {
			names = append(names, columns[i])
		}
	}
	return names
}

func writeFile(names []string, file *os.File) {
	length := len(names)
	fmt.Println(length)
	for i := 0; i < length; i++ {
		file.WriteString(names[i] + "\n")
	}
}
