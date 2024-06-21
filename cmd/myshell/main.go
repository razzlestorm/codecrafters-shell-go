package main

import (
	"bufio"
	"path/filepath"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/commands"
)

var commandlist *commands.CommandHandler

func evaluate(input string, comms *commands.CommandHandler){
	args := strings.Split(input, " ")
	
	command, optional := args[0], args[1:]

	output, ok := comms.Commands[command]
	
	if ok {
		output(optional)		
	} else if input == "" {

	} else if command != "" {

		path := os.Getenv("PATH")
		if len(path) == 0 {
			fmt.Printf("%v: not found\n", command)
			return
		} else {
			dirs := strings.Split(path, ":")
			for _, entry := range dirs {
				exec_path := filepath.Join(entry, command)
				if _, err := os.Stat(exec_path); err == nil {
					cmd := exec.Command(exec_path, strings.Join(optional, " "))
					out, err := cmd.Output()

					if err != nil {
						fmt.Printf("%v returned an error: %v", command, err)
						return
					}
					fmt.Printf("%s\n", strings.Trim(string(out), "\n\r "))
					return
				}
				}
			}
			fmt.Printf("%v: not found\n", command)
			return

	} else {
		fmt.Printf("%v: command not found\n", command)
	}
}



func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	msg := make(chan string, 1)

	commandlist := commands.NewCommandHandler()

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
				evaluate(input, commandlist)
				break post
			}
		}
	}
}
