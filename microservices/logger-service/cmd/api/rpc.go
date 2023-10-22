package main

import (
	"context"
	"log-service/data"
	"time"
)

type RPCServer struct {
}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload JsonPayload, resp *string) error {
	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()

	coll := client.Database("logs").Collection("logs")
	_, err := coll.InsertOne(ctx, data.LogEntryModel{
		Name: payload.Name,
		Data: payload.Data,
	})
	if err != nil {
		return err
	}

	*resp = "Log entry created successfully with RPC: " + payload.Name

	return nil
}
