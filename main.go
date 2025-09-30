package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "repo",
		Short: "A simple CLI tool to manage repositories",
	}

	// `init` command
	var initCmd = &cobra.Command{
		Use:   "init [name]",
		Short: "Initialize a new repository",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			repoName := args[0]
			err := initRepo(repoName)
			if err != nil {
				log.Fatalf("Error initializing repo: %v", err)
			}
			fmt.Printf("Repository '%s' initialized successfully!\n", repoName)
		},
	}

	rootCmd.AddCommand(initCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initRepo(name string) error {
	// Create folder
	err := os.Mkdir(name, 0755)
	if err != nil {
		return err
	}

	// Create hidden .repo folder inside
	repoPath := filepath.Join(name, ".repo")
	err = os.Mkdir(repoPath, 0755)
	if err != nil {
		return err
	}

	// Create config file
	configFile := filepath.Join(repoPath, "config")
	err = os.WriteFile(configFile, []byte("repository = "+name+"\n"), 0644)
	if err != nil {
		return err
	}

	return nil
}
