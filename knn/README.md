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
go run cmd/client/main.go -x 1.5 -y 2.5 -k 3
```

Clean
```bash
make clean
```

To confirm the result, run the test script in python
```bash
python3 test.py -x 1.5 -y 2.5 -k 3
```
