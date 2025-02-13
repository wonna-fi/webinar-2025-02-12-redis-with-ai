// This file contains AI generated code that has not been reviewed by a human

package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// RESPType represents the type of a RESP value
type RESPType int

const (
	// SimpleString represents a RESP Simple String (+)
	SimpleString RESPType = iota
	// Error represents a RESP Error (-)
	Error
	// Integer represents a RESP Integer (:)
	Integer
	// BulkString represents a RESP Bulk String ($)
	BulkString
	// Array represents a RESP Array (*)
	Array
)

// RESPValue represents a RESP protocol value
type RESPValue struct {
	Type   RESPType
	Str    string
	Int    int64
	Array  []RESPValue
	IsNull bool
}

// ParseRESP parses a RESP message from a reader
func ParseRESP(r io.Reader) (RESPValue, int, error) {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}

	// Read the type byte
	typ, err := br.ReadByte()
	if err != nil {
		return RESPValue{}, 0, fmt.Errorf("failed to read type: %w", err)
	}
	bytesRead := 1

	switch typ {
	case '+':
		str, n, err := readLine(br)
		if err != nil {
			return RESPValue{}, bytesRead, fmt.Errorf("failed to read simple string: %w", err)
		}
		bytesRead += n
		return RESPValue{Type: SimpleString, Str: str}, bytesRead, nil

	case '$':
		length, n, err := readInteger(br)
		if err != nil {
			return RESPValue{}, bytesRead, fmt.Errorf("failed to read bulk string length: %w", err)
		}
		bytesRead += n

		if length == -1 {
			return RESPValue{Type: BulkString, IsNull: true}, bytesRead, nil
		}

		str, n, err := readBulkString(br, length)
		if err != nil {
			return RESPValue{}, bytesRead, fmt.Errorf("failed to read bulk string data: %w", err)
		}
		bytesRead += n

		return RESPValue{Type: BulkString, Str: str}, bytesRead, nil

	case '*':
		length, n, err := readInteger(br)
		if err != nil {
			return RESPValue{}, bytesRead, fmt.Errorf("failed to read array length: %w", err)
		}
		bytesRead += n

		if length == -1 {
			return RESPValue{Type: Array, IsNull: true}, bytesRead, nil
		}

		array := make([]RESPValue, length)
		for i := int64(0); i < length; i++ {
			val, n, err := ParseRESP(br)
			if err != nil {
				return RESPValue{}, bytesRead, fmt.Errorf("failed to read array element %d: %w", i, err)
			}
			bytesRead += n
			array[i] = val
		}

		return RESPValue{Type: Array, Array: array}, bytesRead, nil

	default:
		return RESPValue{}, bytesRead, fmt.Errorf("unknown type byte: %c", typ)
	}
}

// readLine reads until \r\n and returns the string without the \r\n
func readLine(r *bufio.Reader) (string, int, error) {
	var bytesRead int
	var line []byte

	for {
		b, err := r.ReadByte()
		if err != nil {
			return "", bytesRead, err
		}
		bytesRead++

		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' && line[len(line)-1] == '\n' {
			return string(line[:len(line)-2]), bytesRead, nil
		}
	}
}

// readInteger reads an integer followed by \r\n
func readInteger(r *bufio.Reader) (int64, int, error) {
	line, n, err := readLine(r)
	if err != nil {
		return 0, n, err
	}
	i, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		return 0, n, fmt.Errorf("invalid integer: %w", err)
	}
	return i, n, nil
}

// readBulkString reads a bulk string of the given length followed by \r\n
func readBulkString(r *bufio.Reader, length int64) (string, int, error) {
	if length < 0 {
		return "", 0, fmt.Errorf("negative bulk string length")
	}

	// Read the string data
	data := make([]byte, length)
	n, err := io.ReadFull(r, data)
	if err != nil {
		return "", n, err
	}

	// Read the \r\n
	cr, err := r.ReadByte()
	if err != nil {
		return "", n, err
	}
	n++

	lf, err := r.ReadByte()
	if err != nil {
		return "", n, err
	}
	n++

	if cr != '\r' || lf != '\n' {
		return "", n, fmt.Errorf("malformed bulk string")
	}

	return string(data), n, nil
}
