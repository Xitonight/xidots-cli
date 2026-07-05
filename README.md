# xidots-cli

> [!WARNING]
> **This project is deprecated and no longer maintained.**

I have since moved to [NixOS](https://nixos.org/), where my system configuration
is declared natively in the Nix language. With that setup there is no longer any
need for a separate dotfiles manager, so `xidots-cli` is no longer used or
developed.

The repository is kept around for reference and history. Feel free to fork it if
it's useful to you, but please don't expect any updates, bug fixes, or support.

---

A CLI tool to install and manage dotfiles, with health checks and an interactive
TUI. Built with Go, [Cobra](https://github.com/spf13/cobra), and
[Bubble Tea](https://github.com/charmbracelet/bubbletea).

## Commands

| Command   | Description                                                        |
| --------- | ------------------------------------------------------------------ |
| `xidots`  | Launch the interactive TUI.                                        |
| `install` | Run the full dotfiles installation pipeline (sync, stow, configure). |
| `health`  | Run health checks on all configured steps.                          |
| `sync`    | Clone or pull the dotfiles repository.                              |
