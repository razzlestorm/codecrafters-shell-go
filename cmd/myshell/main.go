package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"github.com/razzlestorm/codecrafters-shell-go/cmd/myshell/commands"
)

var commandlist *commands.CommandHandler

func evaluate(input string, comms *commands.CommandHandler){
	args := strings.Split(input, " ")
	
	command, optional := args[0], args[1:]

	output, ok := comms.Commands[command]
	
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

	commandlist = commands.NewCommandHandler()

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
				evaluate(input, &commandlist)
				break post
			}
		}
	}
}
