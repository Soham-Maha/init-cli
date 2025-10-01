package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "backend-cli",
	Short: "CLI to generate a Golang backend boilerplate by CodeChef VIT",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, "Error running the CLI \n", err)
		os.Exit(1)
	}
}