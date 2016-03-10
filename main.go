package main

import (
	"github.com/c3mb0/go-do-work"
	"fmt"
	"github.com/forrest321/one-hundred-monkeys/rands"
	"github.com/forrest321/one-hundred-monkeys/files"
)
const (
	FileName = "scrawl.txt"
)
type monkey struct {
	smash  string
	makewords chan string
}

func monkeySmash() string {
	return rands.RandStrings(50)
}

func (a monkey) DoWork() {
	a.smash = monkeySmash()
	a.makewords <- a.smash
}

func main() {
	result := make(chan string)
	monkmonk := monkey{
		smash:  rands.RandStrings(50),
		makewords: result,
	}
	pool := gdw.NewWorkerPool(10)
	defer pool.Close()
	pool.Add(monkmonk, 100)
	go func() {
		for res := range result {
			files.WriteString(FileName, res)
			fmt.Print(res, " ")
		}
	}()
	pool.Wait()
	close(result) // close the result channel after the pool has completed
	fmt.Println()
}