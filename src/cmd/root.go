package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"xidots-cli/src/tui"
)

var rootCmd = &cobra.Command{
	Use:   "xidots",
	Short: "Xidots dotfiles installer and manager",
	Long:  "A CLI tool to install and manage your dotfiles with health checks and an interactive TUI.",
	Run: func(cmd *cobra.Command, args []string) {
		runTUI()
	},
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install dotfiles",
	Long:  "Run the full dotfiles installation pipeline: sync repo, install packages, stow files, and configure services.",
	Run: func(cmd *cobra.Command, args []string) {
		runTUI()
	},
}

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Check dotfiles health",
	Long:  "Run health checks on all configured steps to verify your system is properly set up.",
	Run: func(cmd *cobra.Command, args []string) {
		runTUI()
	},
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync dotfiles repository",
	Long:  "Clone or pull the dotfiles repository to the local directory.",
	Run: func(cmd *cobra.Command, args []string) {
		runTUI()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(healthCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(syncCmd)
}

func runTUI() {
	tui.Run()
}
