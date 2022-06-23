# go-dreamhost

Go API for interacting with DreamHost DNS provider's interface

## Usage

```go
key := "<your dreamhost API key>"
var httpClient http.Client // if you need a configured http client
client, err := api.NewClient(key, httpClient)
```

## Known issues

If the DNS record is created through this API, Dreamhost persists the record with a trailing `.`.
So if you've created a record with `record=test.mydomain.com` value, upon listing and deletion the record will be displayed as `record=test.mydomain.com.` (not the trailing dot).

## Operations

Currently the API supports the following operations:

### Adding DNS record

```go
input := api.DNSRecordInput{}
err := client.AddDNSRecord(context.Background(), input)
```

### Listing DNS records

```go
records, err = client.ListDNSRecords(context.Background())
```

### Removing DNS records

```go
input := api.DNSRecordInput{}
err := client.RemoveDNSRecord(context.Background(), input)
```

## Examples

See `example/`.
