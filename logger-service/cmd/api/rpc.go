package main

import (
	"context"
	"log"
	"log-service/data"
	"time"
)

// RPCServer is the type for RPC Server. Methods that take this as a receriver are available
// over RPC, as long as they are exported
type RPCServer struct{}

// RPCPayload is the type for data received from RPC
type RPCPayload struct {
	Name string
	Data string
}

// LogInfo writes our payload to mongo
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("error writing to mongo", err)
	}

	*resp = "Processed payload via RPC" + payload.Name
	return nil
}
