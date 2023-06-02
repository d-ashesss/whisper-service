package main

import (
	whisperpb "github.com/d-ashesss/whisper-service/proto"
	"google.golang.org/grpc"
)

func main() {
	cfg := NewConfig()
	srv := grpc.NewServer()
	whisperpb.RegisterWhisperServiceServer(srv, &WhisperServiceServer{})
	app := NewApp(cfg, srv)
	app.Run()
}
