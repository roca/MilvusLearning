package main

import (
	"context"
	"fmt"
	"log"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

func main() {

	// retry := client.RetryRateLimitOption{
	// 	MaxRetry:   1,
	// 	MaxBackoff: time.Duration(10 * time.Second), //time.Duration
	// }

	config := client.Config{
		Address: "localhost:19530",
		//RetryRateLimit: &retry,
	}

	// NewGrpcClient
	// milvusClient, err := client.NewGrpcClient(
	// 	context.Background(), // ctx
	// 	"localhost:19530",    // addr
	// )

	// NewClient
	milvusClient, err := client.NewClient(
		context.Background(),
		config,
	)
	if err != nil {
		log.Fatal("failed to connect to Milvus:", err.Error())
	}
	defer closeConnection(milvusClient)

	fmt.Println("connected to Milvus")
}

func closeConnection(milvusClient client.Client) {
	if milvusClient == nil {
		return
	}
	err := milvusClient.Close()
	if err != nil {
		log.Fatal("failed to close Milvus connection:", err.Error())
	}
}
