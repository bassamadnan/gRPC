# Real-time Document - gRPC

First, we would like to thank  Aadhav Vignesh and the other contributors of PairPad --A collaborative text editor
using CRDTs and WebSockets. This had implementations of an interface and the CRDT datastructure, which was what we needed as Go
does not have a CRDT module (or an Operation Transform (OT) module unlike JavaScript (`ot.js`))

Link to PairPad repositry - [Link](https://github.com/burntcarrot/pairpad)
Link to Blog - [Link](https://databases.systems/posts/collaborative-editor)
Link to system design video of google docs (motivaiton behind choosing CRDT) - [Link](https://youtu.be/YCjVIDv0zQY?si=Id-VS1GMsHBhqN8Z)

While we do not use CRDT in the strict sense, we do perform the operation that CRDT describes for addition and insertion. That is to say,
we ONLY use the above project for the interface which is built on top of termbox, and the CRDT data structure implementation.

## Setup
1. To run the code, make sure you delete `logs.txt` (logs are created in appended manner)
2. Build the proto file (these two can be done via the `./toBuild.sh` script)
3. Run `go mod tidy` to install relevant packages.
4. Run the server `go run cmd/server/main.go`
5. Run the client `go run cmd/client/*.go -id 1` (To run a client with id 1)
6. Above command can be used to run multiple clients, run them with different ID's.

Different clients can edit the document at the same time, we do not expect clients to rejoin, though they may crash.
The logs of all the operation can be found in `logs.txt` file
