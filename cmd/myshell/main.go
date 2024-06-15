package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)


var COMMANDS = map[string]func([]string){
	"echo": echo,
	"cd": cd,
	"exit": exit,
}

func cd(input []string){
	if len(input) < 1 {
		fmt.Println("not enough arguments for cd.")
	}
}

func echo(input []string){
	if len(input) < 1 {
		fmt.Println("not enough arguments for echo.")
	}
	fmt.Println(strings.Join(input, " "))
}

func exit(input []string){
	if strings.Join(input, "") != "0" {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func evaluate(input string){
	args := strings.Split(input, " ")
	
	command, optional := args[0], args[1:]

	output, ok := COMMANDS[command]
	
	if ok {
		output(optional)		
	} else if input == "" {

	} else {
		fmt.Printf("%v: command not found\n", command)
	}
}



func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	msg := make(chan string, 1)


	input: for {

		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return
	}
		msg <- input

		post: for {

			select {
			case <-sigs:
				break input

			case s := <-msg:
				input = strings.Trim(s, "\n\r ")
				evaluate(input)
				break post
			}
		}
	}
}
