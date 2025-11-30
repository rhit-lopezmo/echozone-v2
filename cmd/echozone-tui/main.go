package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("usage: echozone-tui <url>")
		os.Exit(1)
	}

	url := args[1]

	cmd := exec.Command("yt-dlp", "-g", "-f", "140", url)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("yt-dlp error:", err)
		os.Exit(1)
	}

	audioURL := strings.TrimSpace(string(out))

	ctx, cancel := context.WithCancel(context.Background())
	cmd = exec.CommandContext(ctx, "mpv", "--no-video", audioURL)
	cmd.Start()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signals
		cancel() // kills MPV
		os.Exit(0)
	}()

	for {

	}
}
