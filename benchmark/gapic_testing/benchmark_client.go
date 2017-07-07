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

package benchmark

import (
	gax "github.com/googleapis/gax-go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	"google.golang.org/grpc"
	grpc_testingpb "google.golang.org/grpc/benchmark/grpc_testing"
)

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	UnaryCall     []gax.CallOption
	StreamingCall []gax.CallOption
}

func defaultClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("testing.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultCallOptions() *CallOptions {
	retry := map[[2]string][]gax.CallOption{}
	return &CallOptions{
		UnaryCall:     retry[[2]string{"default", "non_idempotent"}],
		StreamingCall: retry[[2]string{"default", "non_idempotent"}],
	}
}

// Client is a client for interacting with gRPC Benchmark API.
type Client struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	client grpc_testingpb.BenchmarkServiceClient

	// The call options for this service.
	CallOptions *CallOptions

	// The metadata to be sent with each request.
	xGoogHeader []string
}

// NewClient creates a new benchmark service client.
//
// The benchmark service.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn:        conn,
		CallOptions: defaultCallOptions(),

		client: grpc_testingpb.NewBenchmarkServiceClient(conn),
	}
	c.SetGoogleClientInfo()
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *Client) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.conn.Close()
}

// SetGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) SetGoogleClientInfo(keyval ...string) {
	kv := append([]string{}, keyval...)
	kv = append(kv, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeader = []string{gax.XGoogHeader(kv...)}
}

// UnaryCall one request followed by one response.
// The server returns the client payload as-is.
func (c *Client) UnaryCall(ctx context.Context, req *grpc_testingpb.SimpleRequest, opts ...gax.CallOption) (*grpc_testingpb.SimpleResponse, error) {
	ctx = insertXGoog(ctx, c.xGoogHeader)
	opts = append(c.CallOptions.UnaryCall[0:len(c.CallOptions.UnaryCall):len(c.CallOptions.UnaryCall)], opts...)
	var resp *grpc_testingpb.SimpleResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UnaryCall(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// StreamingCall one request followed by one response.
// The server returns the client payload as-is.
func (c *Client) StreamingCall(ctx context.Context, opts ...gax.CallOption) (grpc_testingpb.BenchmarkService_StreamingCallClient, error) {
	ctx = insertXGoog(ctx, c.xGoogHeader)
	opts = append(c.CallOptions.StreamingCall[0:len(c.CallOptions.StreamingCall):len(c.CallOptions.StreamingCall)], opts...)
	var resp grpc_testingpb.BenchmarkService_StreamingCallClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.StreamingCall(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
