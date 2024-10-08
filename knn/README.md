instructions to run:

build protofiles
```bash
make build_proto
```

build Go binaries

```bash
make build
```

Run the servers
```bash
make make run_servers SERVERS=7
```

Run the client

```bash
go run cmd/client/main.go -x 2.7 -y 5.12 -k 3
```

Clean
```bash
make clean
```
