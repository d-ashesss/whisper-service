package whisper

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type TranscriptionError struct {
	ee *exec.ExitError
}

func (err TranscriptionError) Error() string {
	stderr := bytes.TrimSpace(err.ee.Stderr)
	if err.ee.ProcessState.ExitCode() == 2 || bytes.HasPrefix(stderr, []byte("Traceback")) {
		if idx := bytes.LastIndexByte(stderr, '\n'); idx > 0 {
			stderr = stderr[idx+1:]
		}
	}
	return fmt.Sprintf("whisper: %s\n%s", err.ee, stderr)
}

func (err TranscriptionError) Unwrap() error {
	return err.ee
}

type Service interface {
	Transcribe(ctx context.Context, audiopath string) (string, error)
}

func NewService() Service {
	return &localService{}
}

type localService struct {
}

func (s localService) Transcribe(ctx context.Context, audiopath string) (string, error) {
	args := []string{
		"--output_format", "json",
		"--output_dir", os.TempDir(),
		"--verbose", "False",
		audiopath,
	}
	cmd := exec.CommandContext(ctx, "whisper", args...)

	if _, err := cmd.Output(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return "", TranscriptionError{ee}
		}
		return "", err
	}
	filename := strings.TrimSuffix(path.Base(audiopath), path.Ext(audiopath))
	transcriptpath := os.TempDir() + "/" + filename + ".json"
	log.Println(audiopath)
	log.Println(transcriptpath)
	transcript, err := os.ReadFile(transcriptpath)
	if err != nil {
		return "", fmt.Errorf("read transcription file: %w", err)
	}
	_ = os.Remove(transcriptpath)
	return string(transcript), nil
}
