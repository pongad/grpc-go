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
    grpc_testingpb "google.golang.org/grpc/interop/grpc_testing"
)

import (
    "flag"
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "strings"
    "testing"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/ptypes"
    "golang.org/x/net/context"
    "google.golang.org/api/option"
    status "google.golang.org/genproto/googleapis/rpc/status"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/metadata"
    gstatus "google.golang.org/grpc/status"
)

var _ = io.EOF
var _ = ptypes.MarshalAny
var _ status.Status

type mockTestServer struct {
    // Embed for forward compatibility.
    // Tests will keep working if more methods are added
    // in the future.
    grpc_testingpb.TestServiceServer

    reqs []proto.Message

    // If set, all calls return this error.
    err error

    // responses to return if err == nil
    resps []proto.Message
}

func (s *mockTestServer) EmptyCall(ctx context.Context, req *grpc_testingpb.Empty) (*grpc_testingpb.Empty, error) {
    md, _ := metadata.FromIncomingContext(ctx)
    if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
        return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
    }
    s.reqs = append(s.reqs, req)
    if s.err != nil {
        return nil, s.err
    }
    return s.resps[0].(*grpc_testingpb.Empty), nil
}

func (s *mockTestServer) UnaryCall(ctx context.Context, req *grpc_testingpb.SimpleRequest) (*grpc_testingpb.SimpleResponse, error) {
    md, _ := metadata.FromIncomingContext(ctx)
    if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
        return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
    }
    s.reqs = append(s.reqs, req)
    if s.err != nil {
        return nil, s.err
    }
    return s.resps[0].(*grpc_testingpb.SimpleResponse), nil
}

func (s *mockTestServer) StreamingOutputCall(req *grpc_testingpb.StreamingOutputCallRequest, stream grpc_testingpb.TestService_StreamingOutputCallServer) error {
    md, _ := metadata.FromIncomingContext(stream.Context())
    if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
        return fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
    }
    s.reqs = append(s.reqs, req)
    if s.err != nil {
        return s.err
    }
    for _, v := range s.resps {
        if err := stream.Send(v.(*grpc_testingpb.StreamingOutputCallResponse)); err != nil {
            return err
        }
    }
    return nil
}

func (s *mockTestServer) StreamingInputCall(stream grpc_testingpb.TestService_StreamingInputCallServer) error {
    md, _ := metadata.FromIncomingContext(stream.Context())
    if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
        return fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
    }
    for {
        if req, err := stream.Recv(); err == io.EOF {
            break
        } else if err != nil {
            return err
        } else {
            s.reqs = append(s.reqs, req)
        }
    }
    if s.err != nil {
        return s.err
    }
    return stream.SendAndClose(s.resps[0].(*grpc_testingpb.StreamingInputCallResponse))
}

func (s *mockTestServer) FullDuplexCall(stream grpc_testingpb.TestService_FullDuplexCallServer) error {
    md, _ := metadata.FromIncomingContext(stream.Context())
    if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
        return fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
    }
    for {
        if req, err := stream.Recv(); err == io.EOF {
            break
        } else if err != nil {
            return err
        } else {
            s.reqs = append(s.reqs, req)
        }
    }
    if s.err != nil {
        return s.err
    }
    for _, v := range s.resps {
        if err := stream.Send(v.(*grpc_testingpb.StreamingOutputCallResponse)); err != nil {
            return err
        }
    }
    return nil
}

func (s *mockTestServer) HalfDuplexCall(stream grpc_testingpb.TestService_HalfDuplexCallServer) error {
    md, _ := metadata.FromIncomingContext(stream.Context())
    if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
        return fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
    }
    for {
        if req, err := stream.Recv(); err == io.EOF {
            break
        } else if err != nil {
            return err
        } else {
            s.reqs = append(s.reqs, req)
        }
    }
    if s.err != nil {
        return s.err
    }
    for _, v := range s.resps {
        if err := stream.Send(v.(*grpc_testingpb.StreamingOutputCallResponse)); err != nil {
            return err
        }
    }
    return nil
}


// clientOpt is the option tests should use to connect to the test server.
// It is initialized by TestMain.
var clientOpt option.ClientOption

var (
    mockTest mockTestServer
)

func TestMain(m *testing.M) {
    flag.Parse()

    serv := grpc.NewServer()
    grpc_testingpb.RegisterTestServiceServer(serv, &mockTest)

    lis, err := net.Listen("tcp", "localhost:0")
    if err != nil {
        log.Fatal(err)
    }
    go serv.Serve(lis)

    conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
    }
    clientOpt = option.WithGRPCConn(conn)

    os.Exit(m.Run())
}

func TestTestServiceEmptyCall(t *testing.T) {
    var expectedResponse *grpc_testingpb.Empty = &grpc_testingpb.Empty{}

    mockTest.err = nil
    mockTest.reqs = nil

    mockTest.resps = append(mockTest.resps[:0], expectedResponse)

    var request *grpc_testingpb.Empty = &grpc_testingpb.Empty{}

    c, err := NewTestClient(context.Background(), clientOpt)
    if err != nil {
        t.Fatal(err)
    }

    resp, err := c.EmptyCall(context.Background(), request)

    if err != nil {
        t.Fatal(err)
    }

    if want, got := request, mockTest.reqs[0]; !proto.Equal(want, got) {
        t.Errorf("wrong request %q, want %q", got, want)
    }

    if want, got := expectedResponse, resp; !proto.Equal(want, got) {
        t.Errorf("wrong response %q, want %q)", got, want)
    }
}

func TestTestServiceEmptyCallError(t *testing.T) {
    errCode := codes.PermissionDenied
    mockTest.err = gstatus.Error(errCode, "test error")

    var request *grpc_testingpb.Empty = &grpc_testingpb.Empty{}

    c, err := NewTestClient(context.Background(), clientOpt)
    if err != nil {
        t.Fatal(err)
    }

    resp, err := c.EmptyCall(context.Background(), request)

    if st, ok := gstatus.FromError(err); !ok {
        t.Errorf("got error %v, expected grpc error", err)
    } else if c := st.Code(); c != errCode {
        t.Errorf("got error code %q, want %q", c, errCode)
    }
    _ = resp
}
func TestTestServiceUnaryCall(t *testing.T) {
    var username string = "username-265713450"
    var oauthScope string = "oauthScope443818668"
    var expectedResponse = &grpc_testingpb.SimpleResponse{
        Username: username,
        OauthScope: oauthScope,
    }

    mockTest.err = nil
    mockTest.reqs = nil

    mockTest.resps = append(mockTest.resps[:0], expectedResponse)

    var request *grpc_testingpb.SimpleRequest = &grpc_testingpb.SimpleRequest{}

    c, err := NewTestClient(context.Background(), clientOpt)
    if err != nil {
        t.Fatal(err)
    }

    resp, err := c.UnaryCall(context.Background(), request)

    if err != nil {
        t.Fatal(err)
    }

    if want, got := request, mockTest.reqs[0]; !proto.Equal(want, got) {
        t.Errorf("wrong request %q, want %q", got, want)
    }

    if want, got := expectedResponse, resp; !proto.Equal(want, got) {
        t.Errorf("wrong response %q, want %q)", got, want)
    }
}

func TestTestServiceUnaryCallError(t *testing.T) {
    errCode := codes.PermissionDenied
    mockTest.err = gstatus.Error(errCode, "test error")

    var request *grpc_testingpb.SimpleRequest = &grpc_testingpb.SimpleRequest{}

    c, err := NewTestClient(context.Background(), clientOpt)
    if err != nil {
        t.Fatal(err)
    }

    resp, err := c.UnaryCall(context.Background(), request)

    if st, ok := gstatus.FromError(err); !ok {
        t.Errorf("got error %v, expected grpc error", err)
    } else if c := st.Code(); c != errCode {
        t.Errorf("got error code %q, want %q", c, errCode)
    }
    _ = resp
}
func TestTestServiceStreamingOutputCall(t *testing.T) {
    var expectedResponse *grpc_testingpb.StreamingOutputCallResponse = &grpc_testingpb.StreamingOutputCallResponse{}

    mockTest.err = nil
    mockTest.reqs = nil

    mockTest.resps = append(mockTest.resps[:0], expectedResponse)

    var request *grpc_testingpb.StreamingOutputCallRequest = &grpc_testingpb.StreamingOutputCallRequest{}

    c, err := NewTestClient(context.Background(), clientOpt)
    if err != nil {
        t.Fatal(err)
    }

    stream, err := c.StreamingOutputCall(context.Background(), request)
    if err != nil {
        t.Fatal(err)
    }
    resp, err := stream.Recv()

    if err != nil {
        t.Fatal(err)
    }

    if want, got := request, mockTest.reqs[0]; !proto.Equal(want, got) {
        t.Errorf("wrong request %q, want %q", got, want)
    }

    if want, got := expectedResponse, resp; !proto.Equal(want, got) {
        t.Errorf("wrong response %q, want %q)", got, want)
    }
}

func TestTestServiceStreamingOutputCallError(t *testing.T) {
    errCode := codes.PermissionDenied
    mockTest.err = gstatus.Error(errCode, "test error")

    var request *grpc_testingpb.StreamingOutputCallRequest = &grpc_testingpb.StreamingOutputCallRequest{}

    c, err := NewTestClient(context.Background(), clientOpt)
    if err != nil {
        t.Fatal(err)
    }

    stream, err := c.StreamingOutputCall(context.Background(), request)
    if err != nil {
        t.Fatal(err)
    }
    resp, err := stream.Recv()

    if st, ok := gstatus.FromError(err); !ok {
        t.Errorf("got error %v, expected grpc error", err)
    } else if c := st.Code(); c != errCode {
        t.Errorf("got error code %q, want %q", c, errCode)
    }
    _ = resp
}
func TestTestServiceStreamingInputCall(t *testing.T) {
    var aggregatedPayloadSize int32 = 944910220
    var expectedResponse = &grpc_testingpb.StreamingInputCallResponse{
        AggregatedPayloadSize: aggregatedPayloadSize,
    }

    mockTest.err = nil
    mockTest.reqs = nil

    mockTest.resps = append(mockTest.resps[:0], expectedResponse)

    var request *grpc_testingpb.StreamingInputCallRequest = &grpc_testingpb.StreamingInputCallRequest{}

    c, err := NewTestClient(context.Background(), clientOpt)
    if err != nil {
        t.Fatal(err)
    }

    stream, err := c.StreamingInputCall(context.Background())
    if err != nil {
        t.Fatal(err)
    }
    resp, err := stream.CloseAndRecv()

    if err != nil {
        t.Fatal(err)
    }

    if want, got := request, mockTest.reqs[0]; !proto.Equal(want, got) {
        t.Errorf("wrong request %q, want %q", got, want)
    }

    if want, got := expectedResponse, resp; !proto.Equal(want, got) {
        t.Errorf("wrong response %q, want %q)", got, want)
    }
}

func TestTestServiceStreamingInputCallError(t *testing.T) {
    errCode := codes.PermissionDenied
    mockTest.err = gstatus.Error(errCode, "test error")

    var request *grpc_testingpb.StreamingInputCallRequest = &grpc_testingpb.StreamingInputCallRequest{}

    c, err := NewTestClient(context.Background(), clientOpt)
    if err != nil {
        t.Fatal(err)
    }

    stream, err := c.StreamingInputCall(context.Background())
    if err != nil {
        t.Fatal(err)
    }
    resp, err := stream.CloseAndRecv()

    if st, ok := gstatus.FromError(err); !ok {
        t.Errorf("got error %v, expected grpc error", err)
    } else if c := st.Code(); c != errCode {
        t.Errorf("got error code %q, want %q", c, errCode)
    }
    _ = resp
}
func TestTestServiceFullDuplexCall(t *testing.T) {
    var expectedResponse *grpc_testingpb.StreamingOutputCallResponse = &grpc_testingpb.StreamingOutputCallResponse{}

    mockTest.err = nil
    mockTest.reqs = nil

    mockTest.resps = append(mockTest.resps[:0], expectedResponse)

    var request *grpc_testingpb.StreamingOutputCallRequest = &grpc_testingpb.StreamingOutputCallRequest{}

    c, err := NewTestClient(context.Background(), clientOpt)
    if err != nil {
        t.Fatal(err)
    }

    stream, err := c.FullDuplexCall(context.Background())
    if err != nil {
        t.Fatal(err)
    }
    if err := stream.Send(request); err != nil {
        t.Fatal(err)
    }
    if err := stream.CloseSend(); err != nil {
        t.Fatal(err)
    }
    resp, err := stream.Recv()

    if err != nil {
        t.Fatal(err)
    }

    if want, got := request, mockTest.reqs[0]; !proto.Equal(want, got) {
        t.Errorf("wrong request %q, want %q", got, want)
    }

    if want, got := expectedResponse, resp; !proto.Equal(want, got) {
        t.Errorf("wrong response %q, want %q)", got, want)
    }
}

func TestTestServiceFullDuplexCallError(t *testing.T) {
    errCode := codes.PermissionDenied
    mockTest.err = gstatus.Error(errCode, "test error")

    var request *grpc_testingpb.StreamingOutputCallRequest = &grpc_testingpb.StreamingOutputCallRequest{}

    c, err := NewTestClient(context.Background(), clientOpt)
    if err != nil {
        t.Fatal(err)
    }

    stream, err := c.FullDuplexCall(context.Background())
    if err != nil {
        t.Fatal(err)
    }
    if err := stream.Send(request); err != nil {
        t.Fatal(err)
    }
    if err := stream.CloseSend(); err != nil {
        t.Fatal(err)
    }
    resp, err := stream.Recv()

    if st, ok := gstatus.FromError(err); !ok {
        t.Errorf("got error %v, expected grpc error", err)
    } else if c := st.Code(); c != errCode {
        t.Errorf("got error code %q, want %q", c, errCode)
    }
    _ = resp
}
func TestTestServiceHalfDuplexCall(t *testing.T) {
    var expectedResponse *grpc_testingpb.StreamingOutputCallResponse = &grpc_testingpb.StreamingOutputCallResponse{}

    mockTest.err = nil
    mockTest.reqs = nil

    mockTest.resps = append(mockTest.resps[:0], expectedResponse)

    var request *grpc_testingpb.StreamingOutputCallRequest = &grpc_testingpb.StreamingOutputCallRequest{}

    c, err := NewTestClient(context.Background(), clientOpt)
    if err != nil {
        t.Fatal(err)
    }

    stream, err := c.HalfDuplexCall(context.Background())
    if err != nil {
        t.Fatal(err)
    }
    if err := stream.Send(request); err != nil {
        t.Fatal(err)
    }
    if err := stream.CloseSend(); err != nil {
        t.Fatal(err)
    }
    resp, err := stream.Recv()

    if err != nil {
        t.Fatal(err)
    }

    if want, got := request, mockTest.reqs[0]; !proto.Equal(want, got) {
        t.Errorf("wrong request %q, want %q", got, want)
    }

    if want, got := expectedResponse, resp; !proto.Equal(want, got) {
        t.Errorf("wrong response %q, want %q)", got, want)
    }
}

func TestTestServiceHalfDuplexCallError(t *testing.T) {
    errCode := codes.PermissionDenied
    mockTest.err = gstatus.Error(errCode, "test error")

    var request *grpc_testingpb.StreamingOutputCallRequest = &grpc_testingpb.StreamingOutputCallRequest{}

    c, err := NewTestClient(context.Background(), clientOpt)
    if err != nil {
        t.Fatal(err)
    }

    stream, err := c.HalfDuplexCall(context.Background())
    if err != nil {
        t.Fatal(err)
    }
    if err := stream.Send(request); err != nil {
        t.Fatal(err)
    }
    if err := stream.CloseSend(); err != nil {
        t.Fatal(err)
    }
    resp, err := stream.Recv()

    if st, ok := gstatus.FromError(err); !ok {
        t.Errorf("got error %v, expected grpc error", err)
    } else if c := st.Code(); c != errCode {
        t.Errorf("got error code %q, want %q", c, errCode)
    }
    _ = resp
}