// Copyright 2017, Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package interop

import (
	gax "github.com/googleapis/gax-go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	"google.golang.org/grpc"
	grpc_testingpb "google.golang.org/grpc/interop/grpc_testing"
)

// TestCallOptions contains the retry settings for each method of TestClient.
type TestCallOptions struct {
	EmptyCall           []gax.CallOption
	UnaryCall           []gax.CallOption
	StreamingOutputCall []gax.CallOption
	StreamingInputCall  []gax.CallOption
	FullDuplexCall      []gax.CallOption
	HalfDuplexCall      []gax.CallOption
}

func defaultTestClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("testing.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultTestCallOptions() *TestCallOptions {
	retry := map[[2]string][]gax.CallOption{}
	return &TestCallOptions{
		EmptyCall:           retry[[2]string{"default", "non_idempotent"}],
		UnaryCall:           retry[[2]string{"default", "non_idempotent"}],
		StreamingOutputCall: retry[[2]string{"default", "non_idempotent"}],
		StreamingInputCall:  retry[[2]string{"default", "non_idempotent"}],
		FullDuplexCall:      retry[[2]string{"default", "non_idempotent"}],
		HalfDuplexCall:      retry[[2]string{"default", "non_idempotent"}],
	}
}

// TestClient is a client for interacting with gRPC Testing API.
type TestClient struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	testClient grpc_testingpb.TestServiceClient

	// The call options for this service.
	CallOptions *TestCallOptions

	// The metadata to be sent with each request.
	xGoogHeader []string
}

// NewTestClient creates a new test service client.
//
// A simple service to test the various types of RPCs and experiment with
// performance with various types of payload.
func NewTestClient(ctx context.Context, opts ...option.ClientOption) (*TestClient, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultTestClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &TestClient{
		conn:        conn,
		CallOptions: defaultTestCallOptions(),

		testClient: grpc_testingpb.NewTestServiceClient(conn),
	}
	c.SetGoogleClientInfo()
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *TestClient) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TestClient) Close() error {
	return c.conn.Close()
}

// SetGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TestClient) SetGoogleClientInfo(keyval ...string) {
	kv := append([]string{}, keyval...)
	kv = append(kv, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeader = []string{gax.XGoogHeader(kv...)}
}

// EmptyCall one empty request followed by one empty response.
func (c *TestClient) EmptyCall(ctx context.Context, req *grpc_testingpb.Empty, opts ...gax.CallOption) (*grpc_testingpb.Empty, error) {
	ctx = insertXGoog(ctx, c.xGoogHeader)
	opts = append(c.CallOptions.EmptyCall[0:len(c.CallOptions.EmptyCall):len(c.CallOptions.EmptyCall)], opts...)
	var resp *grpc_testingpb.Empty
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testClient.EmptyCall(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UnaryCall one request followed by one response.
// The server returns the client payload as-is.
func (c *TestClient) UnaryCall(ctx context.Context, req *grpc_testingpb.SimpleRequest, opts ...gax.CallOption) (*grpc_testingpb.SimpleResponse, error) {
	ctx = insertXGoog(ctx, c.xGoogHeader)
	opts = append(c.CallOptions.UnaryCall[0:len(c.CallOptions.UnaryCall):len(c.CallOptions.UnaryCall)], opts...)
	var resp *grpc_testingpb.SimpleResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testClient.UnaryCall(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// StreamingOutputCall one request followed by a sequence of responses (streamed download).
// The server returns the payload with client desired type and sizes.
func (c *TestClient) StreamingOutputCall(ctx context.Context, req *grpc_testingpb.StreamingOutputCallRequest, opts ...gax.CallOption) (grpc_testingpb.TestService_StreamingOutputCallClient, error) {
	ctx = insertXGoog(ctx, c.xGoogHeader)
	opts = append(c.CallOptions.StreamingOutputCall[0:len(c.CallOptions.StreamingOutputCall):len(c.CallOptions.StreamingOutputCall)], opts...)
	var resp grpc_testingpb.TestService_StreamingOutputCallClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testClient.StreamingOutputCall(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// StreamingInputCall a sequence of requests followed by one response (streamed upload).
// The server returns the aggregated size of client payload as the result.
func (c *TestClient) StreamingInputCall(ctx context.Context, opts ...gax.CallOption) (grpc_testingpb.TestService_StreamingInputCallClient, error) {
	ctx = insertXGoog(ctx, c.xGoogHeader)
	opts = append(c.CallOptions.StreamingInputCall[0:len(c.CallOptions.StreamingInputCall):len(c.CallOptions.StreamingInputCall)], opts...)
	var resp grpc_testingpb.TestService_StreamingInputCallClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testClient.StreamingInputCall(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FullDuplexCall a sequence of requests with each request served by the server immediately.
// As one request could lead to multiple responses, this interface
// demonstrates the idea of full duplexing.
func (c *TestClient) FullDuplexCall(ctx context.Context, opts ...gax.CallOption) (grpc_testingpb.TestService_FullDuplexCallClient, error) {
	ctx = insertXGoog(ctx, c.xGoogHeader)
	opts = append(c.CallOptions.FullDuplexCall[0:len(c.CallOptions.FullDuplexCall):len(c.CallOptions.FullDuplexCall)], opts...)
	var resp grpc_testingpb.TestService_FullDuplexCallClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testClient.FullDuplexCall(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// HalfDuplexCall a sequence of requests followed by a sequence of responses.
// The server buffers all the client requests and then serves them in order. A
// stream of responses are returned to the client when the server starts with
// first request.
func (c *TestClient) HalfDuplexCall(ctx context.Context, opts ...gax.CallOption) (grpc_testingpb.TestService_HalfDuplexCallClient, error) {
	ctx = insertXGoog(ctx, c.xGoogHeader)
	opts = append(c.CallOptions.HalfDuplexCall[0:len(c.CallOptions.HalfDuplexCall):len(c.CallOptions.HalfDuplexCall)], opts...)
	var resp grpc_testingpb.TestService_HalfDuplexCallClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testClient.HalfDuplexCall(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
