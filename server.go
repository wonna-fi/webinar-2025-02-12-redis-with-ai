// This file contains AI generated code that has not been reviewed by a human

package main

import (
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
}

// NewServer creates a new Redis-lite server instance
func NewServer() *Server {
	return &Server{
		quit: make(chan struct{}),
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

	// TODO: Implement RESP protocol handling
	// For now, just keep the connection open until the client disconnects
	buffer := make([]byte, 1024)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				log.Printf("Error reading from connection: %v", err)
			}
			return
		}
	}
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
