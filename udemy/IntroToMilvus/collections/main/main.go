package main

import (
	"collections"
	"context"
	"log"
)

func main() {
	defer collections.CloseConnection(collections.MilvusClient)
	if collections.MilvusClient == nil {
		log.Fatal("MilvusClient is nil")
	}

	err := collections.ListAllCollection(context.Background(), collections.MilvusClient)
	if err != nil {
		log.Fatal("failed to list collections:", err.Error())
	}
}
