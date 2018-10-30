package main

import (
	"bufio"
	"fmt"
	"os"

	"./board"
	"./robot"
)

func main() {
	consolescanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")

	archie := robot.NewRobot("Archie")
	board, err := board.NewBoard("5x5")
	if err != nil {
		panic(err)
	}
	archie.LandOn(board)

	for consolescanner.Scan() {
		command := consolescanner.Text()
		err := archie.Execute(command)
		if err != nil {
			fmt.Printf("ERROR: %s\n> ", err)
			continue
		}

		if archie.IsPrintingCommand() {
			fmt.Printf("%s\n> ", archie.PrintPosition())
		} else {
			fmt.Print("> ")
		}
	}

	if err := consolescanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
