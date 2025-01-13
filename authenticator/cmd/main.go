package main

import (
	"context"
	coprocess "github.com/TykTechnologies/tyk-protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

// SimpleDispatcher implements the Tyk gRPC dispatcher interface.
type SimpleDispatcher struct {
}

// Dispatch always responds based on an environment variable "ALLOW_REQUEST".
func (d *SimpleDispatcher) Dispatch(ctx context.Context, object *coprocess.Object) (*coprocess.Object, error) {

	allowRequest := os.Getenv("ALLOW_REQUEST")
	allow, err := strconv.ParseBool(allowRequest)
	if err != nil {
		allow = false // Default to false if parsing fails or env is unset
	}

	if allow == false {
		object.Request.ReturnOverrides.ResponseCode = http.StatusBadRequest
		object.Request.ReturnOverrides.ResponseError = "forbidden"
		return object, nil
	}

	hook := object.GetHookName()
	log.Println(hook, " is called!")

	log.Println("Dispatch called")
	return object, nil
}

func (d *SimpleDispatcher) DispatchEvent(ctx context.Context, event *coprocess.Event) (*coprocess.EventReply, error) {
	log.Println("DispatchEvent called")
	return &coprocess.EventReply{}, nil
}

func main() {
	// Set up gRPC server
	listener, err := net.Listen("tcp", ":9800")
	if err != nil {
		log.Fatalf("Failed to listen on port 9800: %v", err)
	}

	grpcServer := grpc.NewServer()
	dispatcher := &SimpleDispatcher{}

	// Register the dispatcher service
	coprocess.RegisterDispatcherServer(grpcServer, dispatcher)
	log.Println("gRPC server listening on port 9800")

	// Start the server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
