package main

import (
	"context"
	"log"
	"os"
	"vector-db/books"
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

	b := books.Books{CollectionName: "BBooks"}

	// Create a Collection
	err = b.CreateCollection()
	if err != nil {
		log.Println("failed to create Books collection:", err.Error())
		os.Exit(1)
	}

	err = collections.ListAllCollection(context.Background(), collections.MilvusClient)
	if err != nil {
		log.Println("failed to list Books collections:", err.Error())
		os.Exit(1)
	}

	err = collections.RenameCollection(context.Background(), collections.MilvusClient, "BBooks", "books")
	if err != nil {
		log.Println("failed to rename Books collection:", err.Error())
		os.Exit(1)
	}

	err = collections.ListAllCollection(context.Background(), collections.MilvusClient)
	if err != nil {
		log.Println("failed to list Books collections:", err.Error())
		os.Exit(1)
	}

	recordCount, err := b.CreateBooks()
	if err != nil {
		log.Println("failed to create Books:", err.Error())
		os.Exit(1)
	}
	log.Println("Books created:", recordCount)

	err = b.DeleteBooks("book_id in [0, 999]")
	if err != nil {
		log.Println("failed to delete Books:", err.Error())
		os.Exit(1)
	}
	log.Println("1000 Books deleted from original", recordCount)

	err = b.BuildIndex()
	if err != nil {
		log.Println("failed to build index:", err.Error())
		os.Exit(1)
	}
	log.Println("Index built on Books collection")

	err = collections.DropCollection(context.Background(), collections.MilvusClient, "books")
	if err != nil {
		log.Println("failed to drop Books collections:", err.Error())
		os.Exit(1)
	}

	err = collections.ListAllCollection(context.Background(), collections.MilvusClient)
	if err != nil {
		log.Println("failed to list Books collections:", err.Error())
		os.Exit(1)
	}

}
