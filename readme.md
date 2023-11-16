# go-grpc-exercise

- rpc
- message
- repeated field
- error handing
- server
- client
- dial without tls

## Run

### Generate grpc

```bash
make generate
```

### Run server

```bash
go mod download

# grpc server available on :8080
make run-server

make run-client
```