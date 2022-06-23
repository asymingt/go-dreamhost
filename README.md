# go-dreamhost

Go API for interacting with DreamHost DNS provider's interface

## Usage

```go
key := "<your dreamhost API key>"
var httpClient http.Client // if you need a configured http client
client, err := api.NewClient(key, httpClient)
```

## Operations

Currently the API supports the following operations:

### Adding DNS record

```go
input := api.DNSRecordInput{}
err := client.AddDNSRecords(context.Background(), input)
```

### Listing DNS records

```go
records, err = client.ListDNSRecords(context.Background())
```

## Examples

See `example/`.
