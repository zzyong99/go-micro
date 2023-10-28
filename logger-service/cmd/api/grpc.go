package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"

	"google.golang.org/grpc"
)

// LogServer is the type
type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

// WriteLog implements the code generated from protoc
func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	// write log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, nil
	}

	// return response
	res := &logs.LogResponse{Result: "grpc logged!"}
	return res, nil
}

func (app *Config) gRPCListen() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatal("Fail to listen for gRPC (Listener): ", err)
	}

	srv := grpc.NewServer()

	logs.RegisterLogServiceServer(srv, &LogServer{Models: app.Models})

	log.Printf("gRPC server started on port %s", gRpcPort)

	if err := srv.Serve(listen); err != nil {
		log.Fatal("Failed to listen for gRPC (Server): ", err)
	}
}
