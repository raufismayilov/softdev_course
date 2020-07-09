package commandprocessor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Handler is a type of function that handles actions
type Handler func(string, []string) bool

// CommandProcessor can be used for command processing
type CommandProcessor struct {
	m          map[string]Handler
	defHandler Handler
}

// Handle ...
func (p *CommandProcessor) Handle(cmd string, h Handler) {
	p.initializeMap()

	if cmd == "" {
		panic("command cannot be empty")
	}

	if h == nil {
		panic("handler cannot be nil")
	}

	p.m[cmd] = h
}

// HandleDefault ...
func (p *CommandProcessor) HandleDefault(h Handler) {
	p.initializeMap()

	if h == nil {
		panic("handler cannot be nil")
	}

	p.defHandler = h
}

// Run ...
func (p *CommandProcessor) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cmd, args := splitCommand(scanner.Text())

		h, ok := p.m[cmd]

		if !ok {
			if !p.triggerDefault(cmd, args) {
				break
			}
			continue
		}

		if !h(cmd, args) {
			break
		}
	}
}

func (p *CommandProcessor) initializeMap() {
	if p.m == nil {
		p.m = map[string]Handler{}
	}
}

func (p *CommandProcessor) triggerDefault(cmd string, args []string) bool {
	if p.defHandler == nil {
		argsText := strings.Join(args, " ")
		fmt.Printf("unrecognized command: %s, [%s]\n", cmd, argsText)
		return true
	}

	return p.defHandler(cmd, args)
}

func splitCommand(command string) (string, []string) {
	tokens := strings.Split(command, " ")

	if len(tokens) == 0 {
		panic("unexpected empty command")
	}

	return tokens[0], tokens[1:]
}
