package commands

import (
	"fmt"
	"os"
	"strings"
)

	
type CommandHandler struct {
	Commands map[string]func([]string)
} 

func NewCommandHandler() *CommandHandler {
	c := &CommandHandler{
		Commands: make(map[string]func([]string)),
	}
	c.initCommands()
	return c
}

func (c *CommandHandler) initCommands() {
	c.Commands["cd"] = c.cd
	c.Commands["echo"] = c.echo
	c.Commands["exit"] = c.exit
	c.Commands["type"] = c.cmd_type

}


func (c *CommandHandler) cmd_type(input []string) {
	if len(input) < 1 {
		fmt.Println("not enough arguments for type.")
	}

	command := strings.Join(input, "") 
	_, ok := c.Commands[command] 
	if !ok {
		fmt.Printf("%v: not found\n", command)
		return
	}

	fmt.Printf("%v is a shell builtin\n", command)
	return
	}


func (c *CommandHandler) cd(input []string){
	if len(input) < 1 {
		fmt.Println("not enough arguments for cd.")
	}
}

func (c *CommandHandler) echo(input []string){
	if len(input) < 1 {
		fmt.Println("not enough arguments for echo.")
	}
	fmt.Println(strings.Join(input, " "))
}

func (c *CommandHandler) exit(input []string){
	if strings.Join(input, "") != "0" {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}