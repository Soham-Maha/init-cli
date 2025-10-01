package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializing the Project\n",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Initializing the  Project Structure.....\n")
		folders := []string{"cmd/api", "database/queries", "database/schema", "doc", "internal/controller", "internal/middleware", "internal/models", "internal/router", "internal/utils", "internal/db"}
		for _, folder := range folders {
			if err := os.MkdirAll(folder, os.ModePerm); err != nil {
				fmt.Printf("Error making the folders\n")
			}
		}
		files := []string{".env", ".env.example", "DockerFile", "cmd/api/main.go", "internal/controller/controller.go", "internal/router/router.go", "internal/utils/util.go", "internal/middleware/middleware1.go", "internal/models/model1.go", "doc/docs.yaml", "internal/db/db.go", "internal/db/table1.sql.go"}
		for _, file := range files {
			if _, err := os.Create(file); err != nil {
				fmt.Print("Error Creating file\n", file)
			}
		}
		fmt.Printf("Files created successfully\n")
		fmt.Printf("All folders created successfully")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}