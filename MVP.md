# Echozone — MVP Specification
A terminal-based YouTube audio player built in Go using a TUI interface and MPV as the initial playback backend.

## Overview
Echozone MVP is the foundational version of the application. It focuses on:
- YouTube audio playback
- Basic playlist management
- A simple terminal UI
- A reusable core engine decoupled from UI
- MPV as the audio backend
- yt-dlp used to extract direct audio URLs

This MVP intentionally avoids advanced DSP, visualizers, and custom decoding.

---

## Architecture (MVP)

```
cmd/
    echozone-tui/       # executable (main TUI app)
echozone-core/
    player/             # MPV backend (process control)
    youtube/            # yt-dlp wrapper
    playlist/           # playlist storage and CRUD
    model/              # shared structs (Track, Playlist, PlayerState)
    state/              # core engine state and controller
internal/
    ui/                 # TUI components using tview
```

### Layers

### UI Layer (TUI)
- Displays interface (playlist, now playing, input field)
- Sends commands to the core engine
- Receives events (track updated, state changed)

### Core Engine
- Maintains playback state machine
- Selects next track
- Loads playlists from disk
- Interfaces with audio backend

### Audio Backend (MPV)
- Launches mpv as a subprocess
- Plays audio-only streams
- Basic control (play, pause, stop)
- No DSP or PCM extraction in MVP

### YouTube Extraction Wrapper
- Runs `yt-dlp -g -f 140 <url>`
- Returns direct audio URL for playback

---

## MVP Features

### 1. Play a Single YouTube URL
- User pastes a YouTube video URL into the TUI
- Echozone fetches a direct audio URL via yt-dlp
- MPV plays it without video

### 2. Simple TUI Display
Uses tview with:
- Now Playing panel
- URL input field
- Playlist list
- Basic keybindings:

```
Enter = Play
q = Quit
p = Pause/Resume
n = Next Track
a = Add Track to Playlist
```

### 3. Local Playlist Storage
Playlists stored in:

```
~/.echozone/playlists/*.json
```

Example format:

```json
{
  "name": "favorites",
  "items": [
    { "url": "https://youtu.be/abc123", "title": "Track 1" },
    { "url": "https://youtu.be/xyz789", "title": "Track 2" }
  ]
}
```

Supported operations:
- Create playlists
- Add tracks
- Remove tracks
- Play next/previous

### 4. Core Engine State Machine

States:
- Stopped
- Loading
- Playing
- Paused

The engine owns:
- Current track
- Queue handling
- Events sent to UI

### 5. Config Directory

```
~/.echozone/
    playlists/
    config.json
```

---

## Stretch Goals (Not Required for MVP)

Optional but simple additions:
- Title fetching via yt-dlp JSON
- Persist last-opened playlist
- Basic MPV IPC for improved control

---

## Non-Goals (For Now)

These are explicitly not included in the MVP:
- Visualizers
- Audio DSP or EQ
- Custom Opus/AAC decoding
- GUI version
- YouTube playlist importing
- Skins or themes
- Daemon mode

---

## MVP Completion Checklist

### Core
- [ ] yt-dlp wrapper
- [ ] MPV backend
- [ ] Playlist saving/loading
- [ ] Track model

### Engine
- [ ] State machine
- [ ] Queue logic
- [ ] Event callbacks

### TUI
- [ ] URL input field
- [ ] Playlist view
- [ ] Now-playing panel
- [ ] Keyboard shortcuts
- [ ] Error display

### Integration
- [ ] End-to-end playback from URL
- [ ] Playback from playlist
- [ ] Navigation (next/prev)
- [ ] Clean shutdown

---

## Summary

The Echozone MVP delivers:
- YouTube audio playback through MPV
- A simple TUI interface
- A core engine ready for reuse in future GUI or CLI variants
- A lightweight, modular design forming the basis for future phases such as custom DSP, visualizers, and advanced playback backends.
