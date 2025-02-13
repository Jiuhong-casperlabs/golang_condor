package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_condor/utils"
	"net/http"

	"github.com/make-software/casper-go-sdk/v2/casper"
	"github.com/make-software/casper-go-sdk/v2/types"
)

type MyQueryGlobalStateResult struct {
	ApiVersion  string            `json:"api_version"`
	BlockHeader types.BlockHeader `json:"block_header,omitempty"`
	StoredValue types.StoredValue `json:"stored_value"`
	//MerkleProof is a construction created using a merkle trie that allows verification of the associated hashes.
	// MerkleProof json.RawMessage `json:"merkle_proof"`
}

func main() {
	client := casper.NewRPCClient(casper.NewRPCHandler(utils.ENDPOINT, http.DefaultClient))
	blockHash := "8deb1219c2f8dc15d647fe78d5858a75ebec388c696039e8393cbb43f0c4fcad"
	accountKey := "account-hash-76cd901bcfcc531724af6b27a04ab61134d07e6982c8795330352cbd7528f1f7"
	res, err := client.QueryGlobalStateByBlockHash(context.Background(), blockHash, accountKey, nil)
	if err != nil {
		panic(err)
	}

	mystruct := MyQueryGlobalStateResult{ApiVersion: res.ApiVersion, BlockHeader: res.BlockHeader, StoredValue: res.StoredValue}

	globalStateRes, err := json.MarshalIndent(mystruct, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(globalStateRes))

}
