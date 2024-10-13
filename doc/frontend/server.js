const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const PROTO_PATH = "../docs.proto";

// Load protobuf
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});
const docsProto = grpc.loadPackageDefinition(packageDefinition).docs;
console.log(docsProto);

// Create a client instance
const client = new docsProto.DocumentService(
  "localhost:5050",
  grpc.credentials.createInsecure(),
);

// Make the SayHello request
client.SayHello({ name: "client1" }, (error, response) => {
  if (!error) {
    console.log('Greeting:', response.message);
  } else {
    console.error('Error:', error);
  }
});
