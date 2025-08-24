
# Golang-GRPC Example

This repository demonstrates a complete gRPC implementation in Go, featuring unary, server streaming, client streaming, and bidirectional streaming APIs. The project is organized for clarity and modularity, making it easy to understand and extend.

## Features
- **Unary RPC**: Simple request-response interaction.
- **Server Streaming RPC**: Server sends a stream of responses for a single client request.
- **Client Streaming RPC**: Client sends a stream of requests to the server and receives a single response.
- **Bidirectional Streaming RPC**: Both client and server exchange streams of messages.

## Project Structure
```
Raghav_GRPC/
├── client/
│   ├── bidirectionalStream.go
│   ├── clientStram.go
│   ├── main.go
│   ├── serverStreaming.go
│   └── unaryAPI.go
├── proto/
│   ├── greet_grpc.pb.go
│   ├── greet.pb.go
│   └── greet.proto
├── server/
│   ├── bidirectionalStream.go
│   ├── clientstream.go
│   ├── main.go
│   ├── serverStreaming.go
│   └── unaryAPI.go
├── go.mod
└── go.sum
```

## Getting Started

### Prerequisites
- Go 1.18+
- Protocol Buffers compiler (`protoc`)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

### Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/RaghavMoudgil/Golang-GRPC.git
    cd Golang-GRPC/Raghav_GRPC
    ```
2. Install dependencies:
    ```sh
    go mod tidy
    ```
3. Generate gRPC code from proto files:
    ```sh
    protoc --go_out=. --go-grpc_out=. proto/greet.proto
    ```

### Running the Server
1. Start the gRPC server:
    ```sh
    go run server/main.go
    ```

### Running the Client
1. In a separate terminal, run the client:
    ```sh
    go run client/main.go
    ```

## How It Works
- The server implements all gRPC service methods defined in `proto/greet.proto`.
- The client demonstrates usage of all RPC types by calling respective methods.
- Communication is handled using Protocol Buffers for efficient serialization.

## Customization
- Modify `proto/greet.proto` to add new RPC methods or messages.
- Regenerate Go code after changes using the `protoc` command above.

## Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.








