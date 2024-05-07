# Grafana Alloy Remote Config

This repository contains the protobuf definitions and [connect-go](https://github.com/connectrpc/connect-go) generated code for the [Grafana Alloy's](https://github.com/grafana/alloy) remote configuration features.

This code can be used to implement a remote configuration server for Grafana Alloy.

## Updating the protobuf definitions
Make the changes to the relevant `/api/*/*.proto` files.

To build the container that has the required dependencies, run:
```bash
make build-container
```

To generate the code with the new protobuf definitions, run the following command to regenerate the code in a Docker container:
```bash
make buf-generate
```
