package main

import (
	"context"
	"log"

	"github.com/onflow/flow-go/cmd/util/cmd/common"
)

// Note: this AccessAddress will need to be updated with the correct domain for the spork.
// The AccessNodePubKey generally does not change. See the main README.md for details on
// how to lookup the latest address and public key (if it changes).
var access001 = common.FlowClientConfig{
	AccessAddress:    "access-001.mainnet16.nodes.onflow.org:9001",
	AccessNodePubKey: "ba530d6e593947d1dd2d7f8afcf122efac9070043ce7ffdc62b4c4be9899f9a3b7e57c4c975d484386b4c5ad25a2ede097cbd497942a759a6391ba9cf724f6d9",
}

func main() {
	// create the Access API client
	client, err := common.FlowClient(&access001)
	if err != nil {
		log.Fatalf("error creating gRPC client: %v", err)
	}

	// query the Access API using gRPC over TLS
	header, err := client.GetLatestBlockHeader(context.Background(), true)
	if err != nil {
		log.Fatalf("error getting latest block header: %v", err)
	}

	log.Printf("Latest Sealed Block: %d => %s", header.Height, header.ID)
}
