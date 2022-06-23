package main

import (
	"context"

	"github.com/adamantal/go-dreamhost/api"
)

func main() {
	key := "<dreamhost API key>"
	client, err := api.NewClient(key, nil)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	_, err = client.ListDNSRecords(ctx)
	if err != nil {
		panic(err)
	}

	// print the records

	recordInput := api.DNSRecordInput{
		Record: "test.mydomain.com",
		Value:  "<some other>",
		Type:   api.CNAMERecordType,
	}
	err = client.AddDNSRecord(ctx, recordInput)
	if err != nil {
		panic(err)
	}

	err = client.RemoveDNSRecord(ctx, recordInput)
	if err != nil {
		panic(err)
	}
}
