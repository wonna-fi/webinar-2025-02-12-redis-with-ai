// This file contains AI generated code that has not been reviewed by a human

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"sync"
)

const (
	// DefaultPort is the default port that Redis servers listen on
	DefaultPort = 6379
	// DefaultHost is the default host to listen on (all interfaces)
	DefaultHost = "0.0.0.0"
)

// Server represents our Redis-lite server
type Server struct {
	listener net.Listener
	quit     chan struct{}
	wg       sync.WaitGroup
	registry *CommandRegistry
}

// NewServer creates a new Redis-lite server instance
func NewServer() *Server {
	return &Server{
		quit:     make(chan struct{}),
		registry: NewCommandRegistry(),
	}
}

// Start starts the server on the specified host and port
func (s *Server) Start(ctx context.Context) error {
	addr := fmt.Sprintf("%s:%d", DefaultHost, DefaultPort)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	s.listener = listener

	log.Printf("Server listening on %s", addr)

	s.wg.Add(1)
	go s.serve(ctx)

	return nil
}

// serve handles incoming connections
func (s *Server) serve(ctx context.Context) {
	defer s.wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case <-s.quit:
			return
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				select {
				case <-s.quit:
					return
				default:
					log.Printf("Error accepting connection: %v", err)
					continue
				}
			}
			s.wg.Add(1)
			go s.handleConnection(conn)
		}
	}
}

// handleConnection handles a single client connection
func (s *Server) handleConnection(conn net.Conn) {
	defer s.wg.Done()
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	log.Printf("New connection from %s", remoteAddr)

	reader := bufio.NewReader(conn)

	for {
		// Parse the incoming command
		value, _, err := ParseRESP(reader)
		if err != nil {
			if err.Error() != "EOF" {
				log.Printf("Error parsing command from %s: %v", remoteAddr, err)
			}
			return
		}

		log.Printf("Received value from %s: %+v", remoteAddr, value)

		// Handle the command
		response, err := s.handleCommand(value)
		if err != nil {
			log.Printf("Error handling command from %s: %v", remoteAddr, err)
			// Send error response
			errResp := RESPValue{Type: Error, Str: err.Error()}
			if _, err := conn.Write(SerializeRESP(errResp)); err != nil {
				log.Printf("Error sending error response to %s: %v", remoteAddr, err)
				return
			}
			continue
		}

		// Send the response
		if _, err := conn.Write(SerializeRESP(response)); err != nil {
			log.Printf("Error sending response to %s: %v", remoteAddr, err)
			return
		}
	}
}

// handleCommand processes a command and returns a response
func (s *Server) handleCommand(value RESPValue) (RESPValue, error) {
	if value.Type != Array || len(value.Array) == 0 {
		return RESPValue{}, fmt.Errorf("invalid command format")
	}

	// Get the command name from the first array element
	cmdName := value.Array[0]
	if cmdName.Type != BulkString {
		return RESPValue{}, fmt.Errorf("invalid command name format")
	}

	log.Printf("Looking up command: %q", cmdName.Str)

	// Look up the command
	cmd, ok := s.registry.Get(cmdName.Str)
	if !ok {
		return RESPValue{}, fmt.Errorf("unknown command '%s'", cmdName.Str)
	}

	// Execute the command with the arguments
	return cmd.Execute(value.Array[1:])
}

// Stop gracefully shuts down the server
func (s *Server) Stop() error {
	close(s.quit)
	if s.listener != nil {
		if err := s.listener.Close(); err != nil {
			return fmt.Errorf("error closing listener: %w", err)
		}
	}
	s.wg.Wait()
	return nil
}
