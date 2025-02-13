// This file contains AI generated code that has not been reviewed by a human

package main

import (
	"fmt"
	"strings"
)

// Command represents a Redis command handler
type Command interface {
	// Execute executes the command with the given arguments and returns a RESP value
	Execute(args []RESPValue) (RESPValue, error)
}

// CommandFunc is a function type that implements the Command interface
type CommandFunc func(args []RESPValue) (RESPValue, error)

// Execute implements the Command interface for CommandFunc
func (f CommandFunc) Execute(args []RESPValue) (RESPValue, error) {
	return f(args)
}

// CommandRegistry holds all registered commands
type CommandRegistry struct {
	commands map[string]Command
}

// NewCommandRegistry creates a new command registry
func NewCommandRegistry() *CommandRegistry {
	r := &CommandRegistry{
		commands: make(map[string]Command),
	}
	r.registerBuiltinCommands()
	return r
}

// Register registers a command with the given name
func (r *CommandRegistry) Register(name string, cmd Command) {
	r.commands[strings.ToUpper(name)] = cmd
}

// Get returns the command with the given name
func (r *CommandRegistry) Get(name string) (Command, bool) {
	cmd, ok := r.commands[strings.ToUpper(name)]
	return cmd, ok
}

// registerBuiltinCommands registers all builtin commands
func (r *CommandRegistry) registerBuiltinCommands() {
	// PING command
	r.Register("PING", CommandFunc(func(args []RESPValue) (RESPValue, error) {
		switch len(args) {
		case 0:
			return RESPValue{Type: SimpleString, Str: "PONG"}, nil
		case 1:
			return args[0], nil
		default:
			return RESPValue{}, fmt.Errorf("wrong number of arguments for 'ping' command")
		}
	}))

	// ECHO command
	r.Register("ECHO", CommandFunc(func(args []RESPValue) (RESPValue, error) {
		if len(args) != 1 {
			return RESPValue{}, fmt.Errorf("wrong number of arguments for 'echo' command")
		}
		// Convert the response to a bulk string if it's not already
		if args[0].Type != BulkString {
			return RESPValue{Type: BulkString, Str: args[0].Str}, nil
		}
		return args[0], nil
	}))
}

// SerializeRESP converts a RESPValue to its wire format
func SerializeRESP(v RESPValue) []byte {
	switch v.Type {
	case SimpleString:
		return []byte(fmt.Sprintf("+%s\r\n", v.Str))
	case Error:
		return []byte(fmt.Sprintf("-%s\r\n", v.Str))
	case Integer:
		return []byte(fmt.Sprintf(":%d\r\n", v.Int))
	case BulkString:
		if v.IsNull {
			return []byte("$-1\r\n")
		}
		return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(v.Str), v.Str))
	case Array:
		if v.IsNull {
			return []byte("*-1\r\n")
		}
		var b []byte
		b = append(b, []byte(fmt.Sprintf("*%d\r\n", len(v.Array)))...)
		for _, elem := range v.Array {
			b = append(b, SerializeRESP(elem)...)
		}
		return b
	default:
		// This should never happen in practice
		return []byte(fmt.Sprintf("-ERR unknown value type %d\r\n", v.Type))
	}
}
