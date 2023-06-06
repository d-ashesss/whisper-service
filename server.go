package main

import (
	"fmt"
	whisperpb "github.com/d-ashesss/whisper-service/proto"
	"github.com/d-ashesss/whisper-service/whisper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"
)

// WhisperServiceServer is an implementation of Whisper gRPC server.
type WhisperServiceServer struct {
	whisperpb.UnimplementedWhisperServiceServer
	service whisper.Service
}

// NewServer creates a gRPC Service for Whisper.
func NewServer(srv whisper.Service) *WhisperServiceServer {
	return &WhisperServiceServer{service: srv}
}

// Transcribe performes transcription of audio file into text.
func (s WhisperServiceServer) Transcribe(stream whisperpb.WhisperService_TranscribeServer) error {
	log.Printf("[whisper.transcribe] Received transcription request")
	req, file, err := recvTranscribe(stream)
	if err != nil {
		log.Printf("[whisper.transcribe] ERROR: Failed to receive or save the file: %s", err)
		if _, ok := status.FromError(err); ok {
			return err
		}
		return status.Errorf(codes.Unavailable, "unable to receive the file")
	}
	defer func(name string) {
		if err := os.Remove(name); err != nil {
			log.Printf("[whisper.transcribe] ERROR: Failed to delete tmp file: %s", err)
		}
	}(file.Name())

	opts := make([]whisper.Option, 0)
	if req.Format != "" {
		opts = append(opts, whisper.WithFormat(req.Format))
	}

	transcript, err := s.service.Transcribe(stream.Context(), file.Name(), opts...)
	if err != nil {
		log.Printf("[whisper.transcribe] ERROR: Transcription failed: %s", err)
		return status.Errorf(codes.Internal, "transcription failed")
	}
	r := &whisperpb.TranscribeResponse{Transcription: transcript}
	if err := stream.SendAndClose(r); err != nil {
		log.Printf("[whisper.transcribe] ERROR: Failed to respond: %s", err)
		return err
	}
	log.Printf("[whisper.transcribe] Transcription completed")
	return nil
}

// recvTranscribe receives audio contents from request stream and saves it into temporary file.
// Caller must delete the file after it has been processed.
func recvTranscribe(stream whisperpb.WhisperService_TranscribeServer) (*whisperpb.TranscribeRequest, *os.File, error) {
	file, err := os.CreateTemp("", "*.tmp")
	if err != nil {
		return nil, nil, fmt.Errorf("create file: %w", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	var req *whisperpb.TranscribeRequest
	for true {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, fmt.Errorf("read chunk: %w", err)
		}
		if r.GetChunk() == nil {
			return nil, nil, status.Errorf(codes.InvalidArgument, "invalid request input")
		}
		if _, err := file.Write(r.GetChunk()); err != nil {
			return nil, nil, fmt.Errorf("write chunk: %w", err)
		}
		req = r
	}
	if req == nil {
		return nil, nil, status.Errorf(codes.InvalidArgument, "invalid request input")
	}
	req.Chunk = nil
	return req, file, nil
}
