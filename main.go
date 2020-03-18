package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func process(name string) error {
	dat, err := excelize.OpenFile(name)
	check(err)
	newFile, err := os.Create("EmailList.txt")
	check(err)
	//read columns
	names := readColumn(dat)
	//create new file
	writeToFile(names, newFile)
	newFile.Close()
	return err
}

//check for error
func check(e error) {
	if e != nil {
		log.Fatalf("Failed openning file: %s", e)
	}
}

//read column with the project manager names
//and write to slice
func readColumn(file *excelize.File) []string {
	n := 100
	names := make([]string, 0, 50)
	for i := 1; i < n; i++ {
		h, err := file.GetCellValue("Project Storage", fmt.Sprintf("H%d", i))
		check(err)
		if !find(names, h) {
			names = append(names, h)
		}
	}
	return names
}

//write names of slice into a new file
func writeToFile(names []string, newFile *os.File) {
	length := len(names)
	for i := 1; i < length; i++ {
		fullName := strings.Split(names[i], " ")
		var emailAdress string
		if len(fullName) == 2 {
			emailAdress = fullName[0] + "." + fullName[1] + "@prose.one"
		} else if len(fullName) == 3 {
			emailAdress = fullName[0] + "." + fullName[1] + fullName[2] + "@prose.one"
		}
		//checkUmlaut(emailAdress)
		newFile.WriteString(emailAdress + "\r\n")
	}
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

//checking for umlaut and replacing them
//needs to be improved because inputs with umlauts are not UTF-8 encoded
//need to encode to UTF-8 first to replace umlaut on a case basis
func checkUmlaut(fullName string) string {
	fullName = strings.ToValidUTF8(fullName, "ue")
	if strings.ContainsAny(fullName, "ä") {
		fmt.Println("Test")
		strings.Replace(fullName, "ä", "ae", -1)
	} else if strings.ContainsAny(fullName, "ö") {
		fmt.Println("Test")
		strings.Replace(fullName, "ö", "oe", -1)
	} else if strings.ContainsAny(fullName, "ü") {
		fmt.Println("Test")
		strings.Replace(fullName, "ü", "ue", -1)
	}
	return fullName
}
