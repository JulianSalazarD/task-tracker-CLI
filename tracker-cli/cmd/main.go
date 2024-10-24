package main

import "github.com/spf13/cobra"

func main() {

	// // connect to database and make migration
	// connect()
	// makeMigration()

	// main command
	var rootCmd = &cobra.Command{
		Use:   "jbot",
		Short: "task tracker CLI",
	}
	rootCmd.Execute()

}
