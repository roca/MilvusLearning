package main

import (
	"context"
	"log"
	"os"
	"vector-db/book"
	"vector-db/collections"
)

func main() {
	defer collections.CloseConnection(collections.MilvusClient)
	if collections.MilvusClient == nil {
		log.Fatal("MilvusClient is nil")
	}

	// List Collections
	err := collections.ListAllCollection(context.Background(), collections.MilvusClient)
	if err != nil {
		log.Fatal("failed to list collections:", err.Error())
	}

	b := book.Book{}

	// Create a Collection
	err = b.CreateCollection()
	if err != nil {
		log.Println("failed to create Book collection:", err.Error())
		os.Exit(1)
	}

	_ = collections.ListAllCollection(context.Background(), collections.MilvusClient)
}
