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

package benchmark_test

import (
    "io"

    "golang.org/x/net/context"
    "google.golang.org/grpc/benchmark/gapic_testing"
    grpc_testingpb "google.golang.org/grpc/benchmark/grpc_testing"
)

func ExampleNewClient() {
    ctx := context.Background()
    c, err := benchmark.NewClient(ctx)
    if err != nil {
        // TODO: Handle error.
    }
    // TODO: Use client.
    _ = c
}


func ExampleClient_UnaryCall() {
    ctx := context.Background()
    c, err := benchmark.NewClient(ctx)
    if err != nil {
        // TODO: Handle error.
    }

    req := &grpc_testingpb.SimpleRequest{
        // TODO: Fill request struct fields.
    }
    resp, err := c.UnaryCall(ctx, req)
    if err != nil {
        // TODO: Handle error.
    }
    // TODO: Use resp.
    _ = resp
}

func ExampleClient_StreamingCall() {
    ctx := context.Background()
    c, err := benchmark.NewClient(ctx)
    if err != nil {
        // TODO: Handle error.
    }
    stream, err := c.StreamingCall(ctx)
    if err != nil {
        // TODO: Handle error.
    }
    go func() {
        reqs := []*grpc_testingpb.SimpleRequest{
            // TODO: Create requests.
        }
        for _, req := range reqs {
            if err := stream.Send(req); err != nil {
                // TODO: Handle error.
            }
        }
        stream.CloseSend()
    }()
    for {
        resp, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            // TODO: handle error.
        }
        // TODO: Use resp.
        _ = resp
    }
}

