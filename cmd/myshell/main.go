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
	
	if ok {
		output(optional)		
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

	loop: for {

			select {
			case <-sigs:
				break input
			case s := <-msg:
				input = strings.Trim(s, "\n\r ")
				evaluate(input)
				break loop
			}
		}
	}
}
