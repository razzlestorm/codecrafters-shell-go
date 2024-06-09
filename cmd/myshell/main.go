package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"reflect"
)


var COMMANDS = map[string]func([]string){
	"echo": echo,
	"cd": cd,
}

func echo(input []string){
	if len(input) < 1 {
		fmt.Println("not enough arguments for echo.")
	}
}

func cd(input []string){
	if len(input) < 1 {
		fmt.Println("not enough arguments for cd.")
	}
}

func evaluate(input string){
	args := strings.Split(input, " ")
	
	command, optional := args[0], args[1:]

	output, ok := COMMANDS[command]
	
	fmt.Println(ok)
	fmt.Println(command)
	fmt.Println(optional)

	if ok {
		output(optional)		
	} else {
		fmt.Println(command)
		fmt.Printf("%v: command not found\n", command)
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again.")
		return
	}

	input = strings.Trim(input, "\n\r ")
	evaluate(input)
}
