package stream

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const grpcPort = 8888

func TestGRPCStreamReset(t *testing.T) {
	startGRPCServer()
	proxyPort := startHTTP2ReverseProxy(grpcPort)
	t.Logf("started HTTP reverse proxy at port: %v", proxyPort)

	tests := []struct {
		msg  string
		port int
	}{
		{
			"direct",
			grpcPort,
		},
		{
			"via proxy",
			proxyPort,
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			testClient, stop := newGRPCTestClient(t, fmt.Sprintf(":%d", tt.port))
			defer stop()

			_, err := testClient.SetValue(context.Background(), &SetValueRequest{Key: "foo", Value: "hello"})
			st, ok := status.FromError(err)
			require.True(t, ok)
			t.Logf("got grpc code: %v", st.Code())
			t.Logf("got grpc message: %v", st.Message())
		})
	}
}

func startGRPCServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	panicOnError(err)

	grpcServer := grpc.NewServer()
	RegisterTestServer(grpcServer, newGRPCTestServer())
	go grpcServer.Serve(listener)
}

type grpcTestServer struct {
	sync.Mutex
	keyToValue map[string]string
}

func newGRPCTestServer() *grpcTestServer {
	return &grpcTestServer{keyToValue: make(map[string]string)}
}

func (s *grpcTestServer) SetValue(ctx context.Context, request *SetValueRequest) (response *SetValueResponse, err error) {
	s.Lock()
	defer s.Unlock()

	if request.Value == "" {
		delete(s.keyToValue, request.Key)
	} else {
		s.keyToValue[request.Key] = request.Value
	}
	return &SetValueResponse{}, nil
}

func newGRPCTestClient(t *testing.T, addr string) (TestClient, func() error) {
	clientConn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(t, err, "Failed to dial grpc server")
	return NewTestClient(clientConn), clientConn.Close
}
