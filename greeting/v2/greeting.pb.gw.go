// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: greeting/v2/greeting.proto

/*
Package v2 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v2

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage
var _ = metadata.Join

func request_GreetingService_Greeting_0(ctx context.Context, marshaler runtime.Marshaler, client GreetingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GreetingRequest
	var metadata runtime.ServerMetadata

	msg, err := client.Greeting(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_GreetingService_Greeting_0(ctx context.Context, marshaler runtime.Marshaler, server GreetingServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GreetingRequest
	var metadata runtime.ServerMetadata

	msg, err := server.Greeting(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterGreetingServiceHandlerServer registers the http handlers for service GreetingService to "mux".
// UnaryRPC     :call GreetingServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterGreetingServiceHandlerFromEndpoint instead.
func RegisterGreetingServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server GreetingServiceServer) error {

	mux.Handle("GET", pattern_GreetingService_Greeting_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GreetingService_Greeting_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GreetingService_Greeting_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterGreetingServiceHandlerFromEndpoint is same as RegisterGreetingServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterGreetingServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterGreetingServiceHandler(ctx, mux, conn)
}

// RegisterGreetingServiceHandler registers the http handlers for service GreetingService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterGreetingServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterGreetingServiceHandlerClient(ctx, mux, NewGreetingServiceClient(conn))
}

// RegisterGreetingServiceHandlerClient registers the http handlers for service GreetingService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "GreetingServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "GreetingServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "GreetingServiceClient" to call the correct interceptors.
func RegisterGreetingServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client GreetingServiceClient) error {

	mux.Handle("GET", pattern_GreetingService_Greeting_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GreetingService_Greeting_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GreetingService_Greeting_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_GreetingService_Greeting_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"api", "greeting"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_GreetingService_Greeting_0 = runtime.ForwardResponseMessage
)
