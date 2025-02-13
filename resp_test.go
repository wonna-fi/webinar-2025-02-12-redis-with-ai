// This file contains AI generated code that has not been reviewed by a human

package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestRESPSimpleString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		want     RESPValue
		wantErr  bool
		wantRead int
	}{
		{
			name:     "simple string",
			input:    []byte("+OK\r\n"),
			want:     RESPValue{Type: SimpleString, Str: "OK"},
			wantErr:  false,
			wantRead: 5,
		},
		{
			name:     "empty simple string",
			input:    []byte("+\r\n"),
			want:     RESPValue{Type: SimpleString, Str: ""},
			wantErr:  false,
			wantRead: 3,
		},
		{
			name:    "missing \\r\\n",
			input:   []byte("+OK"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bytes.NewReader(tt.input)
			got, n, err := ParseRESP(reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRESP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRESP() got = %v, want %v", got, tt.want)
			}
			if n != tt.wantRead {
				t.Errorf("ParseRESP() read %d bytes, want %d", n, tt.wantRead)
			}
		})
	}
}

func TestRESPBulkString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		want     RESPValue
		wantErr  bool
		wantRead int
	}{
		{
			name:     "bulk string",
			input:    []byte("$3\r\nfoo\r\n"),
			want:     RESPValue{Type: BulkString, Str: "foo"},
			wantErr:  false,
			wantRead: 9,
		},
		{
			name:     "empty bulk string",
			input:    []byte("$0\r\n\r\n"),
			want:     RESPValue{Type: BulkString, Str: ""},
			wantErr:  false,
			wantRead: 6,
		},
		{
			name:     "null bulk string",
			input:    []byte("$-1\r\n"),
			want:     RESPValue{Type: BulkString, IsNull: true},
			wantErr:  false,
			wantRead: 5,
		},
		{
			name:    "invalid length",
			input:   []byte("$abc\r\n"),
			wantErr: true,
		},
		{
			name:    "missing data",
			input:   []byte("$3\r\nfo"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bytes.NewReader(tt.input)
			got, n, err := ParseRESP(reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRESP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRESP() got = %v, want %v", got, tt.want)
			}
			if n != tt.wantRead {
				t.Errorf("ParseRESP() read %d bytes, want %d", n, tt.wantRead)
			}
		})
	}
}

func TestRESPArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		want     RESPValue
		wantErr  bool
		wantRead int
	}{
		{
			name:  "array of bulk strings",
			input: []byte("*2\r\n$4\r\nECHO\r\n$3\r\nfoo\r\n"),
			want: RESPValue{
				Type: Array,
				Array: []RESPValue{
					{Type: BulkString, Str: "ECHO"},
					{Type: BulkString, Str: "foo"},
				},
			},
			wantErr:  false,
			wantRead: 23,
		},
		{
			name:     "empty array",
			input:    []byte("*0\r\n"),
			want:     RESPValue{Type: Array, Array: []RESPValue{}},
			wantErr:  false,
			wantRead: 4,
		},
		{
			name:     "null array",
			input:    []byte("*-1\r\n"),
			want:     RESPValue{Type: Array, IsNull: true},
			wantErr:  false,
			wantRead: 5,
		},
		{
			name:    "invalid length",
			input:   []byte("*abc\r\n"),
			wantErr: true,
		},
		{
			name:    "missing elements",
			input:   []byte("*2\r\n$4\r\nECHO\r\n"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bytes.NewReader(tt.input)
			got, n, err := ParseRESP(reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRESP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRESP() got = %v, want %v", got, tt.want)
			}
			if n != tt.wantRead {
				t.Errorf("ParseRESP() read %d bytes, want %d", n, tt.wantRead)
			}
		})
	}
}
