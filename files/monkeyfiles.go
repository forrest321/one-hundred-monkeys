package files

import (
	"bufio"
	"os"
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
