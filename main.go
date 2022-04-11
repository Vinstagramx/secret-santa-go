package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Person struct {
	name     string
	given    bool
	received bool
}

func initMsg() {
	fmt.Println("--- Welcome to the Secret Santa Generator! ---")
}

func pullName() *Person {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the name of the participant. To exit, type *.")
	fmt.Print("Enter Name: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "*" {
		return nil
	} else {
		return &Person{
			name:     text,
			given:    false,
			received: false,
		}
	}
}

func generatePair(personList []*Person) (*Person, *Person) {
	// two lists: one with received false, one with given false, then randomly generate index for both, then check if names are same
	var notYetGiven, notYetReceived []*Person
	for i := 0; i < len(personList); i++ {
		if !personList[i].given {
			notYetGiven = append(notYetGiven, personList[i])
		}
		if !personList[i].received {
			notYetReceived = append(notYetReceived, personList[i])
		}
	}

	if len(notYetGiven) == 0 && len(notYetReceived) == 0 {
		return nil, nil
	}

	if len(notYetGiven) == 0 || len(notYetReceived) == 0 {
		panic("not enough givers or receivers - should not happen")
	}

	generate := true
	var newGiver, newReceiver *Person
	for generate {
		randomIndex1 := rand.Intn(len(notYetGiven))
		newGiver = notYetGiven[randomIndex1]
		randomIndex2 := rand.Intn(len(notYetReceived))
		newReceiver = notYetReceived[randomIndex2]

		if newGiver != newReceiver {
			generate = false
			newGiver.given = true
			newReceiver.received = true
		}
	}

	// for i := 0; i < len(personList); i++ {
	// 	if personList[i].name == newGiver.name {
	// 		fmt.Println("given")
	// 		personList[i].given = true
	// 	}
	// 	if personList[i].name == newReceiver.name {
	// 		personList[i].received = true
	// 	}
	// }
	return newGiver, newReceiver

}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func main() {
	initMsg()

	var personList []*Person
	stopPull := false
	for !stopPull {
		newPerson := pullName()
		if newPerson == nil {
			stopPull = true
		} else {
			personList = append(personList, newPerson)
			// fmt.Println(newPerson.name)
			fmt.Println("* Saved! *")
		}
	}

	var giversList, receiversList []*Person
	stopGenerate := false
	for !stopGenerate {
		newGiver, newReceiver := generatePair(personList)
		if newGiver == nil && newReceiver == nil {
			stopGenerate = true
		} else {
			giversList = append(giversList, newGiver)
			receiversList = append(receiversList, newReceiver)
		}

		// for _, p := range personList {
		// 	//fmt.Printf("%v\n", *p)
		// }
	}

	for i := 0; i < len(personList); i++ {
		fmt.Printf("%s, you are giving a gift to %s!\n", giversList[i].name, receiversList[i].name)
		time.Sleep(3 * time.Second)
		clearScreen()
	}
	// var names []string
	// names = append(names, "mi")

	// fmt.Println("I always keep my guitar in my car now.")
	// time.Sleep(2 * time.Second)
	// fmt.Println("It's good for traffic jams.")
}
