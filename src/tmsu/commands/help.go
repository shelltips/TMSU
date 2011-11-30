/*
Copyright 2011 Paul Ruane.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
package main
*/

package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

type HelpCommand struct{}

func (this HelpCommand) Name() string {
	return "help"
}

func (this HelpCommand) Summary() string {
	return "lists commands or provides help for a particular command"
}

func (this HelpCommand) Help() string {
	return `  tmsu help          lists commands
  tmsu help COMMAND  shows help for COMMAND

Shows command summary or, when a command is specified, detailed help for that command.`
}

func (this HelpCommand) Exec(args []string) error {
	switch len(args) {
	case 0:
		this.overview()
	default:
		this.commandHelp(args[0])
	}

	return nil
}

// implementation

func (this HelpCommand) overview() {
	fmt.Println("tmsu")
	fmt.Println()

	var maxWidth uint = 0
	commandNames := make([]string, 0, len(commands))
	for commandName, _ := range commands {
		maxWidth = uint(math.Max(float64(maxWidth), float64(len(commandName))))
		commandNames = append(commandNames, commandName)
	}

	sort.Strings(commandNames)

	for _, commandName := range commandNames {
		command, _ := commands[commandName]

		commandSummary := command.Summary()
		if commandSummary == "" {
			continue
		}

		fmt.Printf("  %-"+strconv.Uitoa(maxWidth)+"v  %v\n", command.Name(), commandSummary)
	}
}

func (this HelpCommand) commandHelp(commandName string) {
	command := commands[commandName]
	if command == nil {
		fmt.Printf("No such command '%v'.\n", commandName)
		return
	}

	fmt.Println(command.Help())
}