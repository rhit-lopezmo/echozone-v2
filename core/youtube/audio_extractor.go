package youtube

import (
	"os/exec"
	"strings"
)

func ExtractAudioURL(videoURL string) (string, error) {
	ytCmd := exec.Command("yt-dlp", "-g", "-f", "140", videoURL)
	out, err := ytCmd.Output()
	if err != nil {
		return "", err
	}

	audioURL := strings.TrimSpace(string(out))

	return audioURL, nil
}
