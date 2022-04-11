package main

import (
	// Importing necessary packages
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Person struct {
	// Struct for participant entries.
	name     string
	given    bool
	received bool
}

func initMsg() {
	// Program initialisation message.
	fmt.Println("--- Welcome to the Secret Santa Generator! ---")
}

func pullName() *Person {
	// Takes a name entry from terminal input.
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the name of the participant. To exit, type *.")
	fmt.Print("Enter Name: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "*" {
		return nil // If * entered, then nil returned as an exit flag.
	} else {
		// Name entered is returned as a pointer to a struct, where name field is entered text.
		return &Person{
			name:     text,
			given:    false,
			received: false,
		}
	}
}

func generatePair(personList []*Person) (*Person, *Person) {
	// Generates a pair - one gift-giver and one receiver, which are then returned.
	var notYetGiven, notYetReceived []*Person
	for i := 0; i < len(personList); i++ {
		// Two lists created: one with given flag false and one with received flag false (i.e. names not yet paired).
		if !personList[i].given {
			notYetGiven = append(notYetGiven, personList[i])
		}
		if !personList[i].received {
			notYetReceived = append(notYetReceived, personList[i])
		}
	}

	if len(notYetGiven) == 0 && len(notYetReceived) == 0 {
		// If both lists are empty then nil exit flags are returned - no more pairing needed.
		return nil, nil
	}

	if len(notYetGiven) == 0 || len(notYetReceived) == 0 {
		// Test - if only one list is empty then panic.
		panic("Not enough givers or receivers - should not happen!!")
	}

	// Randomly generates index to generate a giver and a receiver pair out of the unused names.
	generate := true
	var newGiver, newReceiver *Person
	for generate {
		// Index generated until giver and receiver are not equal (preventing duplicates).
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

	return newGiver, newReceiver
}

func clearScreen() {
	// Function to clear terminal console screen. Currently only works for UNIX systems.
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func loadingMessage() {
	// Loading message to allow time for participants.
	fmt.Print("Loading")
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println("First generated pairing will display in 5 seconds!")
	time.Sleep(5 * time.Second)
	clearScreen()
}

func main() {
	// Main body of Secret Santa program.

	initMsg()

	var personList []*Person
	stopPull := false
	for !stopPull {
		newPerson := pullName()
		if newPerson == nil {
			stopPull = true
		} else {
			personList = append(personList, newPerson)
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
	}

	loadingMessage()

	for i := 0; i < len(personList); i++ {
		clearScreen()
		time.Sleep(5 * time.Second)
		fmt.Printf("%s... \n", giversList[i].name)
		fmt.Printf("You are giving a gift to %s!\n", receiversList[i].name)
		time.Sleep(5 * time.Second)
		clearScreen()
	}
}
