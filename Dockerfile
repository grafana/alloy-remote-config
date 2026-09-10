FROM bufbuild/buf:1.43.0 AS buf

FROM golang:1.27.0@sha256:4013ae0f9e7994f8535c58c811f8f863fbed38b72e0d51e6592156f758d66146

# Copy the buf binary from the buf image to the final image
COPY --from=buf /usr/local/bin/buf /usr/local/bin/buf

# Install protoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2

# Install protoc-gen-connect-go
RUN go install connectrpc.com/connect/cmd/protoc-gen-connect-go@v1.17.0

# Set the working directory inside the container
WORKDIR /api

# Copy the files from the current directory into the container
COPY . .

ENTRYPOINT [ "buf" ]

CMD ["generate", "api", "-o", "./api"]
