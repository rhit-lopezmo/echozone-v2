# Bubble Tea TUI Roadmap

1. Set up scaffolding: add Bubble Tea and Lip Gloss dependencies; update `cmd/tui` to boot a minimal model with only a title/status line.

2. Define model basics: create a model struct (focused pane, selection index, playlists slice, status message) plus a centralized keymap (j/k, h/l, gg/G, enter, q). Wire `Update` for key handling and `View` to render placeholder panes.

3. Build playlist storage: add a package to load/save playlists to `~/.local/share/echozone/playlists.json` (configurable). Implement create/list/delete APIs with tests.

4. Lay out panes: render three panes (Library, Playlist, Now Playing/status), add focus highlighting, and handle scrolling for long lists.

5. Add vim motions: implement j/k step, gg/G jump, h/l focus switching, and enter to select/open playlists. Add tests for key handling and bounds.

6. Bridge playback: expose a player interface over mpv IPC (start, stop, status). In Bubble Tea, trigger play on enter and show “Playing…” status; mock playback in tests.

7. Ingest YouTube playlists: add a command to fetch public playlists via yt-dlp (or placeholder parser), return items, and allow saving as local playlists. Surface errors in the status pane.

8. UX polish: add search (`/`), delete (`d`), and reorder (`J/K`) actions. Improve styling with Lip Gloss and ensure layout adapts to terminal width.

9. Resilience: handle missing binaries, IPC timeouts, and storage errors gracefully; add tests for these code paths.

10. Documentation: update `AGENTS.md`, `GOALS.md`, and `README.md` with TUI usage, keymap, and setup notes as features land.
