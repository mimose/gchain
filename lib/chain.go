// chain base
// build chain with commands

package lib

import "errors"

type Chain struct {
	frozen   bool
	commands []Command
}

func NewChain() Chain {
	return Chain{
		frozen:   false,
		commands: make([]Command, 0),
	}
}

func (chain *Chain) AddCommand(commands ...Command) error {
	if chain.frozen {
		return errors.New("the chain is frozen")
	}
	if commands == nil || len(commands) == 0 {
		return errors.New("pls enter the command")
	}
	for _, command := range commands {
		chain.commands = append(chain.commands, command)
	}
	return nil
}

func (chain *Chain) Execute(context *Context) (bool, error) {
	if chain.frozen {
		return false, errors.New("the chain is frozen")
	}
	if context == nil {
		return false, errors.New("pls enter the context")
	}
	if chain.commands == nil || len(chain.commands) == 0 {
		return false, errors.New("pls enter the command first")
	}
	context.frozen = true
	chain.frozen = true
	executeIndex := 0
	executeBreak := false
	var executeError error = nil
	commandsSize := len(chain.commands)
	for executeIndex = 0; executeIndex < commandsSize; executeIndex++ {
		isBreak, err := chain.commands[executeIndex].Execute(context)
		if err != nil {
			executeError = err
			break
		}
		if isBreak {
			executeBreak = true
			break
		}
	}
	return executeBreak, executeError
}
