package player

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"os"
	"os/exec"
	"time"
)

func Stream(audioURL string, ctx context.Context) (*exec.Cmd, error) {
	ipcSocket := "/tmp/echozone.sock"

	if _, err := os.Stat(ipcSocket); !errors.Is(err, os.ErrNotExist) {
		if err := os.Remove(ipcSocket); err != nil {
			return nil, err
		}
	}

	mpvCmd := exec.CommandContext(
		ctx,
		"mpv",
		"--no-video",
		"--quiet",
		"--idle=yes",
		"--input-ipc-server="+ipcSocket,
		audioURL,
	)

	if err := mpvCmd.Start(); err != nil {
		return nil, err
	}

	return mpvCmd, nil
}

type Player struct {
	conn    net.Conn
	decoder *json.Decoder
}

type PlayerEvent struct {
	Event  string `json:"event"`
	Reason string `json:"reason,omitempty"`
	Name   string `json:"name,omitempty"`
	Data   any    `json:"data,omitempty"`
}

type PlayerCommand struct {
	Command []any `json:"command"`
}

func NewPlayer(socketPath string) (*Player, error) {
	var conn net.Conn
	var err error
	tryCount := 0
	maxTries := 15

	for {
		conn, err = net.Dial("unix", socketPath)
		tryCount++
		if err == nil {
			break
		}

		if tryCount >= maxTries {
			return nil, err
		}

		time.Sleep(time.Millisecond * 15)
	}

	player := Player{
		conn:    conn,
		decoder: json.NewDecoder(conn),
	}

	return &player, nil
}

func (player *Player) Send(pComm PlayerCommand) error {
	if player.conn == nil {
		return errors.New("player is not connected")
	}

	bytes, err := json.Marshal(pComm)
	if err != nil {
		return err
	}

	bytes = append(bytes, '\n')

	_, err = player.conn.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func (player *Player) Close() error {
	return player.conn.Close()
}

func (player *Player) Events() <-chan PlayerEvent {
	events := make(chan PlayerEvent, 10)

	go func() {
		for {
			var event PlayerEvent

			err := player.decoder.Decode(&event)
			if err != nil {
				close(events)
				return
			}

			if event.Event != "" {
				events <- event
			}
		}
	}()

	return events
}
