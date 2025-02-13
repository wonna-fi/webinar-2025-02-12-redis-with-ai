// This file contains AI generated code that has not been reviewed by a human

package main

import (
	"context"
	"testing"
)

func TestNewServer(t *testing.T) {
	server := NewServer()
	if server == nil {
		t.Fatal("NewServer() returned nil")
	}
	if server.quit == nil {
		t.Error("server.quit channel was not initialized")
	}
}

func TestServerStartStop(t *testing.T) {
	server := NewServer()
	ctx := context.Background()

	// Start the server
	if err := server.Start(ctx); err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	// Verify the server is listening
	if server.listener == nil {
		t.Fatal("Server listener was not initialized")
	}

	// Stop the server
	if err := server.Stop(); err != nil {
		t.Errorf("Failed to stop server: %v", err)
	}
}
