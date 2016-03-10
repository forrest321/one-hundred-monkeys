package main

import (
	"github.com/c3mb0/go-do-work"
	"fmt"
	"github.com/forrest321/one-hundred-monkeys/rands"
	"github.com/forrest321/one-hundred-monkeys/files"
	"time"
)

const (
	FileName = "scrawl.txt"
	StringToFind = "Roll Tide"
	TimeLimitInSeconds = 15
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
	var didWeDoIt = false
	startTime := time.Now()
	for didWeDoIt == false {
		goMyPretties()
		gotDid := monkeyfiles.WeDidIt(FileName, StringToFind)
		fmt.Println("we did it?: ", gotDid)
		didWeDoIt = gotDid

		rightNow := time.Now().Sub(startTime).Nanoseconds() / 1000
		if rightNow >= TimeLimitInSeconds {
			fmt.Println("we didn't do it")
		}

	}
	endTime := time.Now()
	fmt.Print("something happened! Started at %s and ended at %s", startTime, endTime)
}

func goMyPretties() {
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
			monkeyfiles.WriteString(FileName, res)
		}
	}()
	pool.Wait()
	close(result) // close the result channel after the pool has completed
}