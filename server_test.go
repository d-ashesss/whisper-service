package main

import (
	"context"
	"errors"
	"github.com/d-ashesss/whisper-service/mocks"
	whisperpb "github.com/d-ashesss/whisper-service/proto"
	"github.com/d-ashesss/whisper-service/whisper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"io"
	"log"
	"net"
	"testing"
)

func server(t *testing.T, service whisper.Service) (whisperpb.WhisperServiceClient, func()) {
	t.Helper()
	log.SetOutput(io.Discard)
	bufsize := 1024 * 1024 * 16
	lis := bufconn.Listen(bufsize)

	server := grpc.NewServer()
	whisperpb.RegisterWhisperServiceServer(server, NewServer(service))
	go func() {
		if err := server.Serve(lis); err != nil {
			t.Errorf("Failed to start the server: %s", err)
		}
	}()
	closer := func() {
		server.Stop()
	}

	conn, err := grpc.Dial("", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to open connection: %s", err)
	}
	client := whisperpb.NewWhisperServiceClient(conn)
	return client, closer
}

func TestWhisperServiceServer_Transcribe(t *testing.T) {
	t.Run("no request data sent", func(t *testing.T) {
		client, closer := server(t, nil)
		defer closer()

		stream, _ := client.Transcribe(context.Background())
		_, err := stream.CloseAndRecv()
		stat := status.Convert(err)
		require.Equalf(t, codes.InvalidArgument, stat.Code(), "expecting status %s, got: %s", codes.InvalidArgument, stat)
	})

	t.Run("nil chunk", func(t *testing.T) {
		client, closer := server(t, nil)
		defer closer()

		stream, _ := client.Transcribe(context.Background())
		err := stream.Send(&whisperpb.TranscribeRequest{Chunk: nil})
		require.NoError(t, err, "failed to send the request")
		_, err = stream.CloseAndRecv()
		stat := status.Convert(err)
		require.Equalf(t, codes.InvalidArgument, stat.Code(), "expecting status %s, got: %s", codes.InvalidArgument, stat)
	})

	t.Run("transcription failed", func(t *testing.T) {
		service := mocks.NewService(t)
		service.On("Transcribe", mock.Anything, mock.Anything).Return("", errors.New("transcription error"))
		client, closer := server(t, service)
		defer closer()

		stream, _ := client.Transcribe(context.Background())
		err := stream.Send(&whisperpb.TranscribeRequest{Chunk: []byte("line1")})
		require.NoError(t, err, "failed to send the request")
		_, err = stream.CloseAndRecv()
		stat := status.Convert(err)
		require.Equalf(t, codes.Internal, stat.Code(), "expecting status %s, got: %s", codes.Internal, stat)
	})

	t.Run("successful request", func(t *testing.T) {
		service := mocks.NewService(t)
		service.On("Transcribe", mock.Anything, mock.Anything).Return("transcribed test", nil)
		client, closer := server(t, service)
		defer closer()

		stream, _ := client.Transcribe(context.Background())
		err := stream.Send(&whisperpb.TranscribeRequest{Chunk: []byte("line1")})
		require.NoError(t, err, "failed to send first line")
		err = stream.Send(&whisperpb.TranscribeRequest{Chunk: []byte("line2")})
		require.NoError(t, err, "failed to send second line")
		res, err := stream.CloseAndRecv()
		require.NoErrorf(t, err, "got unexpected error: %s", err)
		assert.Equal(t, "transcribed test", res.Transcription)
	})
}
