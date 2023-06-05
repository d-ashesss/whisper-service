package main

import (
	whisperpb "github.com/d-ashesss/whisper-service/proto"
	"github.com/d-ashesss/whisper-service/whisper"
	"google.golang.org/grpc"
)

func main() {
	cfg := NewConfig()
	srv := grpc.NewServer()
	service := whisper.NewService()
	whisperpb.RegisterWhisperServiceServer(srv, NewServer(service))
	app := NewApp(cfg, srv)
	app.Run()
}
