# Project Goals (Trackable)

- [ ] Core playback: keep audio extraction (yt-dlp) and playback (mpv IPC) decoupled so both TUI and future GUI reuse the same engine.

- [ ] TUI first: build a Bubble Tea interface with vim-style motions (j/k navigation, h/l pane focus, gg/G jump) for playlist browsing and controls.

- [ ] Playlists: add local storage for named playlists that preserves order and links; support create, rename, reorder, delete, and add/remove tracks.

- [ ] YouTube playlists: ingest public YouTube playlists, display entries cleanly, and allow merging or saving into local playlists.

- [ ] UI roadmap: design panes for queue, library, and playback details; aim for parity-ready GUI once TUI patterns are stable.

- [ ] Reliability: handle missing binaries, network errors, and IPC failures gracefully; add tests around command construction and playlist persistence.

Progress notes:
- Update checkboxes as goals land; add short dated bullets under each item if partial milestones are useful.
