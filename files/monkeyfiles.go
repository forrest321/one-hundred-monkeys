package monkeyfiles

import (
	"bufio"
	"os"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteString(filename, content string) {
	f, err := os.Create(filename)
	check(err)
	w := bufio.NewWriter(f)
	_, err = w.WriteString(content)
	w.Flush()
}

func WeDidIt(fileName, stringToFind string) bool {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		//fmt.Print("err in reading file. %s", err)
		return false
	}

	fileString := string(fileBytes)
	return strings.Contains(fileString, stringToFind)
}