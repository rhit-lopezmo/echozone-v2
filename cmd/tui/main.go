package main

import (
	"context"
	"echozone-v2/core/player"
	"echozone-v2/core/youtube"
	"fmt"
	"os"
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
	audioURL, err := youtube.ExtractAudioURL(videoURL)
	if err != nil {
		fmt.Println("could not extract audio URL:", err)
		os.Exit(1)
	}

	fmt.Println("Starting playback...")

	// Step 2: Run MPV with context so it dies if our program dies
	ctx, cancel := context.WithCancel(context.Background())
	mpvCmd, err := player.Stream(audioURL, ctx)
	if err != nil {
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
