# Stage 1: Use a Golang image to get protoc-gen-go
# The version of the go image should match the version of go in this project
FROM golang:1.22.1 AS go_builder

# Install protoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0

# Stage 2: Use the bufbuild/buf image as base
FROM bufbuild/buf:latest

# Copy protoc-gen-go binary from the go_builder stage to the final image
COPY --from=go_builder /go/bin/protoc-gen-go /usr/bin/protoc-gen-go

# Set the working directory inside the container
WORKDIR /api

# Copy the files from the current directory into the container
COPY . .

ENTRYPOINT [ "buf" ]

CMD ["generate", "api", "-o", "./api"]
