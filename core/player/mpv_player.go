package player

import (
	"context"
	"os/exec"
)

func Stream(audioURL string, ctx context.Context) (*exec.Cmd, error) {
	mpvCmd := exec.CommandContext(ctx, "mpv", "--no-video", "--quiet", audioURL)

	if err := mpvCmd.Start(); err != nil {
		return nil, err
	}

	return mpvCmd, nil
}
