package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// Create the cobra root command
var (
	rootCmd = &cobra.Command{
		Use:   "tail-json",
		Short: "Root command for tail-json",
		Long:  "Root command for tail-json",
	}

	subcommand = &cobra.Command{
		Use:   "logs",
		Short: "Tails a file and formats the contents as JSON",
		Long:  "Tails a file and formats the contents as JSON",
		// execute the `run` function when `subcommand` is called
		Run: run,
	}
)

func main() {
	// Register the `logs` subcommand
	rootCmd.AddCommand(subcommand)

	// Call the cobra command
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func run(_ *cobra.Command, args []string) {
	// Validate the arg exists
	if len(args) == 0 {
		fmt.Println("No file name given!")

		os.Exit(1)
	}

	// Create the tail command string
	tailCmdStr := fmt.Sprintf("tail -f %s | jq -R 'fromjson?'", args[0])

	// Create the tail command
	tailCmd := exec.Command("zsh", "-c", tailCmdStr)

	// Redirect tail command stdout, stderr to current stdout, stderr
	tailCmd.Stdout, tailCmd.Stderr = os.Stdout, os.Stderr

	// Execute the tail command
	err := tailCmd.Run()
	if err != nil {
		fmt.Printf("Tail command failed - %s\n", err)

		os.Exit(1)
	}
}
