// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: collector/v1/collector.proto

package collectorv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/grafana/alloy-remote-config/api/gen/proto/go/collector/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// CollectorServiceName is the fully-qualified name of the CollectorService service.
	CollectorServiceName = "collector.v1.CollectorService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CollectorServiceGetConfigProcedure is the fully-qualified name of the CollectorService's
	// GetConfig RPC.
	CollectorServiceGetConfigProcedure = "/collector.v1.CollectorService/GetConfig"
	// CollectorServiceRegisterCollectorProcedure is the fully-qualified name of the CollectorService's
	// RegisterCollector RPC.
	CollectorServiceRegisterCollectorProcedure = "/collector.v1.CollectorService/RegisterCollector"
	// CollectorServiceUnregisterCollectorProcedure is the fully-qualified name of the
	// CollectorService's UnregisterCollector RPC.
	CollectorServiceUnregisterCollectorProcedure = "/collector.v1.CollectorService/UnregisterCollector"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	collectorServiceServiceDescriptor                   = v1.File_collector_v1_collector_proto.Services().ByName("CollectorService")
	collectorServiceGetConfigMethodDescriptor           = collectorServiceServiceDescriptor.Methods().ByName("GetConfig")
	collectorServiceRegisterCollectorMethodDescriptor   = collectorServiceServiceDescriptor.Methods().ByName("RegisterCollector")
	collectorServiceUnregisterCollectorMethodDescriptor = collectorServiceServiceDescriptor.Methods().ByName("UnregisterCollector")
)

// CollectorServiceClient is a client for the collector.v1.CollectorService service.
type CollectorServiceClient interface {
	// GetConfig returns the collector's configuration.
	GetConfig(context.Context, *connect.Request[v1.GetConfigRequest]) (*connect.Response[v1.GetConfigResponse], error)
	// RegisterCollector registers the collector with the given ID and attributes. It will
	// update the collector's attributes if the collector is already registered and the
	// attributes are different.
	RegisterCollector(context.Context, *connect.Request[v1.RegisterCollectorRequest]) (*connect.Response[v1.RegisterCollectorResponse], error)
	// UnregisterCollector unregisters the collector with the given ID.
	UnregisterCollector(context.Context, *connect.Request[v1.UnregisterCollectorRequest]) (*connect.Response[v1.UnregisterCollectorResponse], error)
}

// NewCollectorServiceClient constructs a client for the collector.v1.CollectorService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCollectorServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) CollectorServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &collectorServiceClient{
		getConfig: connect.NewClient[v1.GetConfigRequest, v1.GetConfigResponse](
			httpClient,
			baseURL+CollectorServiceGetConfigProcedure,
			connect.WithSchema(collectorServiceGetConfigMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		registerCollector: connect.NewClient[v1.RegisterCollectorRequest, v1.RegisterCollectorResponse](
			httpClient,
			baseURL+CollectorServiceRegisterCollectorProcedure,
			connect.WithSchema(collectorServiceRegisterCollectorMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyIdempotent),
			connect.WithClientOptions(opts...),
		),
		unregisterCollector: connect.NewClient[v1.UnregisterCollectorRequest, v1.UnregisterCollectorResponse](
			httpClient,
			baseURL+CollectorServiceUnregisterCollectorProcedure,
			connect.WithSchema(collectorServiceUnregisterCollectorMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// collectorServiceClient implements CollectorServiceClient.
type collectorServiceClient struct {
	getConfig           *connect.Client[v1.GetConfigRequest, v1.GetConfigResponse]
	registerCollector   *connect.Client[v1.RegisterCollectorRequest, v1.RegisterCollectorResponse]
	unregisterCollector *connect.Client[v1.UnregisterCollectorRequest, v1.UnregisterCollectorResponse]
}

// GetConfig calls collector.v1.CollectorService.GetConfig.
func (c *collectorServiceClient) GetConfig(ctx context.Context, req *connect.Request[v1.GetConfigRequest]) (*connect.Response[v1.GetConfigResponse], error) {
	return c.getConfig.CallUnary(ctx, req)
}

// RegisterCollector calls collector.v1.CollectorService.RegisterCollector.
func (c *collectorServiceClient) RegisterCollector(ctx context.Context, req *connect.Request[v1.RegisterCollectorRequest]) (*connect.Response[v1.RegisterCollectorResponse], error) {
	return c.registerCollector.CallUnary(ctx, req)
}

// UnregisterCollector calls collector.v1.CollectorService.UnregisterCollector.
func (c *collectorServiceClient) UnregisterCollector(ctx context.Context, req *connect.Request[v1.UnregisterCollectorRequest]) (*connect.Response[v1.UnregisterCollectorResponse], error) {
	return c.unregisterCollector.CallUnary(ctx, req)
}

// CollectorServiceHandler is an implementation of the collector.v1.CollectorService service.
type CollectorServiceHandler interface {
	// GetConfig returns the collector's configuration.
	GetConfig(context.Context, *connect.Request[v1.GetConfigRequest]) (*connect.Response[v1.GetConfigResponse], error)
	// RegisterCollector registers the collector with the given ID and attributes. It will
	// update the collector's attributes if the collector is already registered and the
	// attributes are different.
	RegisterCollector(context.Context, *connect.Request[v1.RegisterCollectorRequest]) (*connect.Response[v1.RegisterCollectorResponse], error)
	// UnregisterCollector unregisters the collector with the given ID.
	UnregisterCollector(context.Context, *connect.Request[v1.UnregisterCollectorRequest]) (*connect.Response[v1.UnregisterCollectorResponse], error)
}

// NewCollectorServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCollectorServiceHandler(svc CollectorServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	collectorServiceGetConfigHandler := connect.NewUnaryHandler(
		CollectorServiceGetConfigProcedure,
		svc.GetConfig,
		connect.WithSchema(collectorServiceGetConfigMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	collectorServiceRegisterCollectorHandler := connect.NewUnaryHandler(
		CollectorServiceRegisterCollectorProcedure,
		svc.RegisterCollector,
		connect.WithSchema(collectorServiceRegisterCollectorMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
		connect.WithHandlerOptions(opts...),
	)
	collectorServiceUnregisterCollectorHandler := connect.NewUnaryHandler(
		CollectorServiceUnregisterCollectorProcedure,
		svc.UnregisterCollector,
		connect.WithSchema(collectorServiceUnregisterCollectorMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/collector.v1.CollectorService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CollectorServiceGetConfigProcedure:
			collectorServiceGetConfigHandler.ServeHTTP(w, r)
		case CollectorServiceRegisterCollectorProcedure:
			collectorServiceRegisterCollectorHandler.ServeHTTP(w, r)
		case CollectorServiceUnregisterCollectorProcedure:
			collectorServiceUnregisterCollectorHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCollectorServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCollectorServiceHandler struct{}

func (UnimplementedCollectorServiceHandler) GetConfig(context.Context, *connect.Request[v1.GetConfigRequest]) (*connect.Response[v1.GetConfigResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("collector.v1.CollectorService.GetConfig is not implemented"))
}

func (UnimplementedCollectorServiceHandler) RegisterCollector(context.Context, *connect.Request[v1.RegisterCollectorRequest]) (*connect.Response[v1.RegisterCollectorResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("collector.v1.CollectorService.RegisterCollector is not implemented"))
}

func (UnimplementedCollectorServiceHandler) UnregisterCollector(context.Context, *connect.Request[v1.UnregisterCollectorRequest]) (*connect.Response[v1.UnregisterCollectorResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("collector.v1.CollectorService.UnregisterCollector is not implemented"))
}
