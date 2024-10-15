# Instructions
Make sure you have Go installed and run `go mod tidy`
## 1. run the state and loadbalancer first
```bash
go run cmd/state/main.go
go run cmd/loadbalance/main.go
```
## 2. connect servers
```bash
go run cmd/server/main.go -id 0
go run cmd/server/main.go -id 1
go run cmd/server/main.go -id 2
..
```
These start listening on ports 5050 , we do not handle cases of ports more than 9 servers (due to sprintf'ing the last digit on spin up)

## 3. connect riders and drivers

Connect the drivers first (reccomended to connect them all at the start for avoiding unexpected behaviour)
```bash
go run cmd/driver/main.go -id 0
go run cmd/driver/main.go -id 1
go run cmd/driver/main.go -id 2
..
```
Then finally run the riders

```bash
go run cmd/rider/main.go -id 0
go run cmd/rider/main.go -id 1
go run cmd/rider/main.go -id 2
..
```
This will send a request to the server with client details (its id, its pickup and destination). You can check for ride details on the client
by pressing `1` . From here you can respond to the request on the driver side as well.

The prompts to perform operations do appear on the client side, while logs can be found on Server terminals.
