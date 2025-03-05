package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/asymingt/go-dreamhost/api"
)

func main() {
	key := os.Getenv("DREAMHOST_API_KEY")
	if key == "" {
		panic("DREAMHOST_API_KEY environment variable is not set")
	}
	client, err := api.NewClient(key, nil)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	records, err := client.ListDNSRecords(ctx)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(records, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))

	recordInput := api.DNSRecordInput{
		Record: "test.snrio.com",
		Value:  "nandrene.synology.me.",
		Type:   api.CNAMERecordType,
	}
	err = client.RemoveDNSRecord(ctx, recordInput)
	if err != nil {
		panic(err)
	}
	err = client.AddDNSRecord(ctx, recordInput)
	if err != nil {
		panic(err)
	}
}
