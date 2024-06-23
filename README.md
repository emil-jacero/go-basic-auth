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
docker build -t auth-service .
```

### Run the Container

To run the container with the required environment variables, use:

```sh
docker run -e AUTH_USERNAME=testuser -e AUTH_PASSWORD=testpassword -p 9001:9001 auth-service
```

Replace testuser and testpassword with your desired credentials.

### Testing the Service

#### Without Authorization Header

To test the service without an Authorization header, run:

```sh
curl -v http://localhost:9001/
```

You should receive a 403 Forbidden response.

#### With Authorization Header

To test the service with a valid Authorization header, run:

```sh
curl -v -H "Authorization: Bearer testuser:testpassword" http://localhost:9001/
```

You should receive a 200 OK response if the credentials are valid.

#### Example Test Script

You can use the following script to automate testing:

```sh
#!/bin/bash

echo "Testing without Authorization header..."
response=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:9001/)
if [ "$response" -eq 403 ]; then
    echo "PASS: Received 403 Forbidden without Authorization header"
else
    echo "FAIL: Expected 403 Forbidden but received $response"
fi

echo "Testing with valid Authorization header..."
response=$(curl -s -o /dev/null -w "%{http_code}" -H "Authorization: Bearer testuser:testpassword" http://localhost:9001/)
if [ "$response" -eq 200 ]; then
    echo "PASS: Received 200 OK with valid Authorization header"
else
    echo "FAIL: Expected 200 OK but received $response"
fi

echo "Testing with invalid Authorization header..."
response=$(curl -s -o /dev/null -w "%{http_code}" -H "Authorization: Bearer invalidtoken" http://localhost:9001/)
if [ "$response" -eq 403 ]; then
    echo "PASS: Received 403 Forbidden with invalid Authorization header"
else
    echo "FAIL: Expected 403 Forbidden but received $response"
fi
```

Save this script to a file (e.g., test.sh), make it executable, and run it:

```sh
chmod +x test.sh
./test.sh
```

## Code Structure

- main.go: Entry point for the application.
- pkg/auth/v3/auth.go: Contains the logic for handling authorization requests.

## Dependencies

- Envoy Proxy Go control plane: github.com/envoyproxy/go-control-plane
- gRPC: google.golang.org/grpc
- Protocol buffers and wrappers: github.com/golang/protobuf, google.golang.org/genproto/googleapis/rpc
