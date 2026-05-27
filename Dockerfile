FROM bufbuild/buf:1.43.0@sha256:8ded6090dbcf06c56bfe769b6d524687ba8a8cc912c38efe341ea2a503b378e0 AS buf

FROM golang:1.23.2@sha256:ad5c126b5cf501a8caef751a243bb717ec204ab1aa56dc41dc11be089fafcb4f

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
