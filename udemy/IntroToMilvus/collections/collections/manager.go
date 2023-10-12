package collections

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

var MilvusClient *client.Client

func init() {

	ch := make(chan string, 1)

	// Create a context with a timeout of 5 seconds
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// Start the ConnectToMilvus function
	go ConnectToMilvus(ctxTimeout, ch)

	select {
	case <-ctxTimeout.Done():
		fmt.Printf("Context cancelled: %v\n", ctxTimeout.Err())
	case result := <-ch:
		fmt.Printf("Received: %s\n", result)
	}

}

func ConnectToMilvus(ctx context.Context, ch chan string) {
	// NewGrpcClient
	milvusClient, err := client.NewGrpcClient(
		ctx,               // ctx
		"localhost:19530", // addr
	)
	if err != nil {
		ch <- fmt.Sprint("failed to connect to Milvus:", err.Error())
	}

	MilvusClient = &milvusClient
	ch <- "connected to Milvus"
}

func CloseConnection(milvusClient *client.Client) {
	if milvusClient == nil {
		return
	}
	c := *milvusClient
	err := c.Close()
	if err != nil {
		log.Fatal("failed to close Milvus connection:", err.Error())
	}
}

func ListAllCollection(ctx context.Context, client *client.Client) error {
	collections, err := (*client).ListCollections(ctx)
	if err != nil {
		return err
	}
	fmt.Println("collections:", collections)
	return nil
}
