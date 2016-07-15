//
// cmd.go
// Copyright (C) 2016 Ryan Vlaming <ryanvlaming@icloud.com>
//
// Distributed under terms of the MIT license.
//

package janitor

import (
	"fmt"
	"gopkg.in/readline.v1"
	"os"
	"strings"
)

type command struct {
	flags     []string
	must_flag bool
	must_arg  bool
}

var commands = map[string]command{
	"clean":   {[]string{"-noignore", "-defaultdir", "revertlast"}, false, false},
	"install": {[]string{"-dtemplate"}, false, false},
	"quit":    {[]string{}, false, false},
}

func Repl() {
	janitor := NewJanitor()
	rl, err := readline.New("janitor> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		uline, err := rl.Readline()
		if err != nil {
			panic(err)
		}

		janitor.handleArgs(uline)
	}

}

func (j *janitor) handleArgs(rli string) {
	args := strings.Split(rli, " ")
	if val, ok := commands[args[0]]; ok {
		if !val.must_flag && !val.must_arg && len(args) <= 1 {
			fmt.Println("not enough arguments given for command : ", args[0])
		} else {
			//at this point we know the command exists and is argvalidated

			switch args[0] {

			case "clean":
				j.CleanDir(args)

			case "install":

			case "quit":
				os.Exit(1)
			}
		}
	} else {
		fmt.Println("unrecognised command : '", args[0], "'")
	}
}

func FlagGiven(flag string, args []string) bool {
	for i := 0; i < len(args); i++ {
		if args[i] == flag {
			return true
		}
	}
	return false
}

func isValidFlag(command string, flag string) bool {
	if isValidCommand(command) {
		for i := 0; i < len(commands[command].flags); i++ {
			if flag == commands[command].flags[i] {
				return true
			}
		}
	}
	return false
}

func isValidCommand(command string) bool {
	if _, ok := commands[command]; ok {
		return true
	}
	return false
}
