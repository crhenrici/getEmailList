package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func process(name string) error {
	dat, err := os.Open(name)
	check(err)
	newFile, err := os.Create("EmailList.txt")
	check(err)
	readFile(dat, newFile)
	dat.Close()
	//create new file
	newFile.Close()

	return err
}

//check for error
func check(e error) {
	if e != nil {
		log.Fatalf("Failed openning file: %s", e)
	}
}

//read file input
func readFile(file *os.File, newFile *os.File) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textLines []string
	var columns []string
	names := make([]string, 0, 100)
	i := 0
	for scanner.Scan() {
		textLines = append(textLines, scanner.Text())
	}
	for i = 0; i < len(textLines); i++ {
		columns = append(strings.Split(textLines[i], "	"))
		//names as parameter so values don't get lost
		names = getNames(names, columns, newFile)
	}
}

//get the names from the while
//more specifically get the name from a specific column in the file
//slice has to be given as parameter to continue to work with the values from last call
func getNames(names, columns []string, newFile *os.File) []string {
	length := len(columns)

	//iterate through columns
	for i := 0; i < length; i++ {
		//7th column is the name required
		if (i % 7) == 0 {
			if columns[i] != "" {
				trimmedName := strings.TrimSpace(columns[i])
				if !find(names, trimmedName) {
					names = append(names, trimmedName)
					fullName := strings.Split(trimmedName, " ")
					var emailAdress string
					if len(fullName) == 2 {
						emailAdress = fullName[0] + "." + fullName[1] + "@prose.one"
					} else if len(fullName) == 3 {
						emailAdress = fullName[0] + "." + fullName[1] + fullName[2] + "@prose.one"
					}
					newFile.WriteString(emailAdress + "\r\n")
				}
			}
		}
	}
	columns = nil
	//return slice so that values of the slice don't get lost
	return names
}

//search if slice contains given value
func find(slice []string, val string) bool {
	for _, item := range slice {
		if val == item {
			return true
		}
	}
	return false
}
