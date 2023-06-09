package whisper

import (
	"context"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

// runWhisper executes Whisper command and returns path to the file containing the transcription.
func runWhisper(ctx context.Context, audiopath string, opts options) (string, error) {
	args := []string{
		audiopath,
		"--output_format", opts.Format,
		"--output_dir", os.TempDir(),
		"--verbose", "False",
		"--word_timestamps", "True",
	}
	if opts.InitialPrompt != "" {
		args = append(args, "--initial_prompt", opts.InitialPrompt)
	}
	if opts.Language != "" {
		args = append(args, "--language", opts.Language)
	}
	if opts.MaxLineCount > 0 {
		args = append(args, "--max_line_count", strconv.FormatUint(opts.MaxLineCount, 10))
	}
	if opts.MaxLineWidth > 0 {
		args = append(args, "--max_line_width", strconv.FormatUint(opts.MaxLineWidth, 10))
	}
	if opts.Translate {
		args = append(args, "--task", "translate")
	}
	cmd := exec.CommandContext(ctx, "whisper", args...)

	if _, err := cmd.Output(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return "", TranscriptionError{ee}
		}
		return "", err
	}
	filename := strings.TrimSuffix(path.Base(audiopath), path.Ext(audiopath))
	return path.Join(os.TempDir(), filename+"."+opts.Format), nil
}
