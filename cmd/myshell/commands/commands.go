package commands

import (
	"path/filepath"
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
	c.Commands["pwd"] = c.pwd

}


func (c *CommandHandler) pwd(input []string) {
	if len(input) > 1 {
		fmt.Println("pwd expects no arguments, but some were passed.")
		return
	}
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
}


func (c *CommandHandler) cmd_type(input []string) {
	if len(input) < 1 {
		fmt.Println("not enough arguments for type.")
		return
	}

	command := strings.Join(input, "") 
	// built-in first
	_, ok := c.Commands[command] 
	if !ok {
		// Then search path
		path := os.Getenv("PATH")
		if len(path) == 0 {
			fmt.Printf("%v: not found\n", command)
			return
		} else {
			dirs := strings.Split(path, ":")
			for _, entry := range dirs {
				exec_path := filepath.Join(entry, command)
				if _, err := os.Stat(exec_path); err == nil {
					fmt.Printf("%v is %v/%v\n", command, entry, command)
					return
				}
				}
			}
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
