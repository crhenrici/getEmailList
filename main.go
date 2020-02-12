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
	newFile, err := os.Create("EmailList.txt")
	check(err)
	readFile(dat, newFile)
	dat.Close()
	//create new file
	newFile.Close()
	fmt.Println("New File Created!")
	var finish string
	fmt.Println("Press any key to finish")
	fmt.Scanf("%s", &finish)

}

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
		names = getNames(names, columns, newFile)
	}
}

//get the names from the while
//more specifically get the name from a specific column in the file
func getNames(names, columns []string, newFile *os.File) []string {
	length := len(columns)

	for i := 0; i < length; i++ {
		//7th column is the name required
		if (i % 7) == 0 {
			if columns[i] != "" {
				trimmedName := strings.TrimSpace(columns[i])
				if !find(names, trimmedName) {
					//fmt.Println(names)
					names = append(names, trimmedName)
					fullName := strings.Split(trimmedName, " ")
					emailAdress := fullName[0] + "." + fullName[1] + "@prose.one"
					//fmt.Println(emailAdress)
					newFile.WriteString(emailAdress + "\r\n")
				}
			}
		}
	}
	columns = nil
	return names
}

//write names to new file
func writeFile(names []string, file *os.File) {
	//length := len(names)
	//fmt.Println(length)
	/*for i := 0; i < length; i++ {
		file.WriteString(names[i] + "\n")
	} */
}

func find(slice []string, val string) bool {
	for _, item := range slice {
		//fmt.Println(val + " =" + item)
		if val == item {
			return true
		}
	}
	return false
}
