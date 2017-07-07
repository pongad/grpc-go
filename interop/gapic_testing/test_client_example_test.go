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

package interop_test

import (
    "io"

    "golang.org/x/net/context"
    "google.golang.org/grpc/interop/gapic_testing"
    grpc_testingpb "google.golang.org/grpc/interop/grpc_testing"
)

func ExampleNewTestClient() {
    ctx := context.Background()
    c, err := interop.NewTestClient(ctx)
    if err != nil {
        // TODO: Handle error.
    }
    // TODO: Use client.
    _ = c
}


func ExampleTestClient_EmptyCall() {
    ctx := context.Background()
    c, err := interop.NewTestClient(ctx)
    if err != nil {
        // TODO: Handle error.
    }

    req := &grpc_testingpb.Empty{
        // TODO: Fill request struct fields.
    }
    resp, err := c.EmptyCall(ctx, req)
    if err != nil {
        // TODO: Handle error.
    }
    // TODO: Use resp.
    _ = resp
}

func ExampleTestClient_UnaryCall() {
    ctx := context.Background()
    c, err := interop.NewTestClient(ctx)
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

func ExampleTestClient_StreamingOutputCall() {
    ctx := context.Background()
    c, err := interop.NewTestClient(ctx)
    if err != nil {
        // TODO: Handle error.
    }

    req := &grpc_testingpb.StreamingOutputCallRequest{
        // TODO: Fill request struct fields.
    }
    stream, err := c.StreamingOutputCall(ctx, req)
    if err != nil {
        // TODO: Handle error.
    }
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

func ExampleTestClient_StreamingInputCall() {
    ctx := context.Background()
    c, err := interop.NewTestClient(ctx)
    if err != nil {
        // TODO: Handle error.
    }
    stream, err := c.StreamingInputCall(ctx)
    if err != nil {
        // TODO: Handle error.
    }
    reqs := []*grpc_testingpb.StreamingInputCallRequest{
        // TODO: Create requests.
    }
    for _, req := range reqs {
        if err := stream.Send(req); err != nil {
            // TODO: Handle error.
        }
    }
    resp, err := stream.CloseAndRecv()
    if err != nil {
        // TODO: Handle error.
    }
    // TODO: Use resp.
    _ = resp
}

func ExampleTestClient_FullDuplexCall() {
    ctx := context.Background()
    c, err := interop.NewTestClient(ctx)
    if err != nil {
        // TODO: Handle error.
    }
    stream, err := c.FullDuplexCall(ctx)
    if err != nil {
        // TODO: Handle error.
    }
    go func() {
        reqs := []*grpc_testingpb.StreamingOutputCallRequest{
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

func ExampleTestClient_HalfDuplexCall() {
    ctx := context.Background()
    c, err := interop.NewTestClient(ctx)
    if err != nil {
        // TODO: Handle error.
    }
    stream, err := c.HalfDuplexCall(ctx)
    if err != nil {
        // TODO: Handle error.
    }
    go func() {
        reqs := []*grpc_testingpb.StreamingOutputCallRequest{
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

