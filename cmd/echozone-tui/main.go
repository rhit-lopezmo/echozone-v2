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

	videoURL := args[1]

	// Step 1: Extract direct audio URL using yt-dlp
	ytCmd := exec.Command("yt-dlp", "-g", "-f", "140", videoURL)
	out, err := ytCmd.Output()
	if err != nil {
		fmt.Println("yt-dlp error:", err)
		os.Exit(1)
	}

	audioURL := strings.TrimSpace(string(out))
	fmt.Println("Starting playback...")

	// Step 2: Run MPV with context so it dies if our program dies
	ctx, cancel := context.WithCancel(context.Background())
	mpvCmd := exec.CommandContext(ctx, "mpv", "--no-video", "--quiet", audioURL)

	if err := mpvCmd.Start(); err != nil {
		fmt.Println("mpv start error:", err)
		os.Exit(1)
	}

	// Step 3: Handle Ctrl+C (SIGINT/SIGTERM)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signals
		fmt.Println("\nStopping playback...")

		cancel() // stop command context
		if mpvCmd.Process != nil {
			mpvCmd.Process.Kill() // ensure MPV is gone
		}

		os.Exit(0)
	}()

	// Step 4: Wait for MPV to exit (normal playback end)
	err = mpvCmd.Wait()
	if err != nil {
		// It's OK if MPV was killed by us
		if !strings.Contains(err.Error(), "killed") {
			fmt.Println("mpv exited with error:", err)
		}
	}

	fmt.Println("Playback finished.")
}
