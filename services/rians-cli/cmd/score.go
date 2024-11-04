package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// scoreCmd represents the score command
var scoreCmd = &cobra.Command{
	Use:   "score",
	Short: "Run the score-k8s CLI from your custom CLI",
	Long:  `This subcommand wraps the score-k8s CLI and forwards commands to it.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create the command to call score-k8s
		scoreCmd := exec.Command("score-k8s", args...)
		scoreCmd.Stdout = os.Stdout
		scoreCmd.Stderr = os.Stderr

		// Run the command and handle errors
		if err := scoreCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error running score-k8s: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(scoreCmd)
}
