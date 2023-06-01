package main

import (
	whisperpb "github.com/d-ashesss/whisper-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type WhisperServiceServer struct {
	whisperpb.UnimplementedWhisperServiceServer
}

func (WhisperServiceServer) Transcribe(whisperpb.WhisperService_TranscribeServer) error {
	log.Printf("[whisper.transcribe] Received transcription request")
	return status.Errorf(codes.Unimplemented, "method Transcribe not implemented")
}
