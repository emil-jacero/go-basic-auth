# Basic Auth Service with Envoy Proxy External Authorization

This project implements a basic authentication service to be used with Envoy Proxy's external authorization HTTP filter. The service is written in Go and uses HTTP. It checks the `Authorization` header of incoming requests and validates it against a single set of credentials provided via environment variables.

## Prerequisites

- Go 1.22
- Docker

## Environment Variables

The service requires the following environment variables to be set:

- `AUTH_USERNAME`: The username for authentication.
- `AUTH_PASSWORD`: The password for authentication.

## Building and Running the Service

### Build the Docker Image

To build the Docker image, run:

```sh
make docker-build
```

### Run the Container

To run the container with the required environment variables, use:

```sh
docker run -e AUTH_USERNAME=testuser -e AUTH_PASSWORD=testpassword -p 9001:9001 grpc-basic-auth
```

Replace testuser and testpassword with your desired credentials.

### Testing the Service

#### Without Authorization Header

To test the service without an Authorization header, run:

```sh

```

## Code Structure

- main.go: Entry point for the application.
- pkg/auth/v3/auth.go: Contains the logic for handling authorization requests.

## Dependencies

- Envoy Proxy Go control plane: github.com/envoyproxy/go-control-plane
- gRPC: google.golang.org/grpc
- Protocol buffers and wrappers: github.com/golang/protobuf, google.golang.org/genproto/googleapis/rpc
