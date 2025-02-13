// This file contains AI generated code that has not been reviewed by a human

package main

import (
	"reflect"
	"testing"
)

func TestPingCommand(t *testing.T) {
	registry := NewCommandRegistry()
	cmd, ok := registry.Get("PING")
	if !ok {
		t.Fatal("PING command not found in registry")
	}

	tests := []struct {
		name    string
		args    []RESPValue
		want    RESPValue
		wantErr bool
	}{
		{
			name: "no arguments",
			args: []RESPValue{},
			want: RESPValue{Type: SimpleString, Str: "PONG"},
		},
		{
			name: "one argument",
			args: []RESPValue{{Type: BulkString, Str: "hello"}},
			want: RESPValue{Type: BulkString, Str: "hello"},
		},
		{
			name:    "too many arguments",
			args:    []RESPValue{{Type: BulkString, Str: "hello"}, {Type: BulkString, Str: "world"}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cmd.Execute(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEchoCommand(t *testing.T) {
	registry := NewCommandRegistry()
	cmd, ok := registry.Get("ECHO")
	if !ok {
		t.Fatal("ECHO command not found in registry")
	}

	tests := []struct {
		name    string
		args    []RESPValue
		want    RESPValue
		wantErr bool
	}{
		{
			name:    "no arguments",
			args:    []RESPValue{},
			wantErr: true,
		},
		{
			name: "one argument",
			args: []RESPValue{{Type: BulkString, Str: "hello"}},
			want: RESPValue{Type: BulkString, Str: "hello"},
		},
		{
			name: "one argument with spaces",
			args: []RESPValue{{Type: BulkString, Str: "hello world"}},
			want: RESPValue{Type: BulkString, Str: "hello world"},
		},
		{
			name:    "too many arguments",
			args:    []RESPValue{{Type: BulkString, Str: "hello"}, {Type: BulkString, Str: "world"}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cmd.Execute(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSetCommands(t *testing.T) {
	registry := NewCommandRegistry()

	// Get SET command
	setCmd, ok := registry.Get("SET")
	if !ok {
		t.Fatal("SET command not found in registry")
	}

	// Get GET command
	getCmd, ok := registry.Get("GET")
	if !ok {
		t.Fatal("GET command not found in registry")
	}

	tests := []struct {
		name    string
		cmd     Command
		args    []RESPValue
		want    RESPValue
		wantErr bool
	}{
		{
			name: "set key value",
			cmd:  setCmd,
			args: []RESPValue{
				{Type: BulkString, Str: "mykey"},
				{Type: BulkString, Str: "myvalue"},
			},
			want: RESPValue{Type: SimpleString, Str: "OK"},
		},
		{
			name: "get existing key",
			cmd:  getCmd,
			args: []RESPValue{
				{Type: BulkString, Str: "mykey"},
			},
			want: RESPValue{Type: BulkString, Str: "myvalue"},
		},
		{
			name: "get non-existent key",
			cmd:  getCmd,
			args: []RESPValue{
				{Type: BulkString, Str: "nonexistent"},
			},
			want: RESPValue{Type: BulkString, IsNull: true},
		},
		{
			name:    "set without value",
			cmd:     setCmd,
			args:    []RESPValue{{Type: BulkString, Str: "mykey"}},
			wantErr: true,
		},
		{
			name:    "set with invalid key type",
			cmd:     setCmd,
			args:    []RESPValue{{Type: SimpleString, Str: "mykey"}, {Type: BulkString, Str: "myvalue"}},
			wantErr: true,
		},
		{
			name:    "set with invalid value type",
			cmd:     setCmd,
			args:    []RESPValue{{Type: BulkString, Str: "mykey"}, {Type: SimpleString, Str: "myvalue"}},
			wantErr: true,
		},
		{
			name:    "get without key",
			cmd:     getCmd,
			args:    []RESPValue{},
			wantErr: true,
		},
		{
			name:    "get with invalid key type",
			cmd:     getCmd,
			args:    []RESPValue{{Type: SimpleString, Str: "mykey"}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cmd.Execute(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSerializeRESP(t *testing.T) {
	tests := []struct {
		name string
		v    RESPValue
		want []byte
	}{
		{
			name: "simple string",
			v:    RESPValue{Type: SimpleString, Str: "OK"},
			want: []byte("+OK\r\n"),
		},
		{
			name: "error",
			v:    RESPValue{Type: Error, Str: "Error message"},
			want: []byte("-Error message\r\n"),
		},
		{
			name: "integer",
			v:    RESPValue{Type: Integer, Int: 42},
			want: []byte(":42\r\n"),
		},
		{
			name: "bulk string",
			v:    RESPValue{Type: BulkString, Str: "hello"},
			want: []byte("$5\r\nhello\r\n"),
		},
		{
			name: "null bulk string",
			v:    RESPValue{Type: BulkString, IsNull: true},
			want: []byte("$-1\r\n"),
		},
		{
			name: "array",
			v: RESPValue{
				Type: Array,
				Array: []RESPValue{
					{Type: BulkString, Str: "PING"},
					{Type: BulkString, Str: "hello"},
				},
			},
			want: []byte("*2\r\n$4\r\nPING\r\n$5\r\nhello\r\n"),
		},
		{
			name: "null array",
			v:    RESPValue{Type: Array, IsNull: true},
			want: []byte("*-1\r\n"),
		},
		{
			name: "empty array",
			v:    RESPValue{Type: Array, Array: []RESPValue{}},
			want: []byte("*0\r\n"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SerializeRESP(tt.v)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SerializeRESP() = %q, want %q", got, tt.want)
			}
		})
	}
}
