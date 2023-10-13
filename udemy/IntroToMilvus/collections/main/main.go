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

	b := book.Book{CollectionName: "BBook"}

	// Create a Collection
	err = b.CreateCollection()
	if err != nil {
		log.Println("failed to create Book collection:", err.Error())
		os.Exit(1)
	}

	err = collections.ListAllCollection(context.Background(), collections.MilvusClient)
	if err != nil {
		log.Println("failed to list Book collections:", err.Error())
		os.Exit(1)
	}

	err = collections.RenameCollection(context.Background(), collections.MilvusClient, "BBook", "book")
	if err != nil {
		log.Println("failed to rename Book collection:", err.Error())
		os.Exit(1)
	}

	err = collections.ListAllCollection(context.Background(), collections.MilvusClient)
	if err != nil {
		log.Println("failed to list Book collections:", err.Error())
		os.Exit(1)
	}

	err = collections.DropCollection(context.Background(), collections.MilvusClient, "book")
	if err != nil {
		log.Println("failed to drop Book collections:", err.Error())
		os.Exit(1)
	}

	err = collections.ListAllCollection(context.Background(), collections.MilvusClient)
	if err != nil {
		log.Println("failed to list Book collections:", err.Error())
		os.Exit(1)
	}

}
