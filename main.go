package main

import (
	whisperpb "github.com/d-ashesss/whisper-service/proto"
	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	whisperpb.RegisterWhisperServiceServer(srv, &WhisperServiceServer{})
	app := NewApp(srv)
	app.Run()
}
