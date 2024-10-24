package main

import (
	"tracker-cli/internal/service"
	"tracker-cli/internal/storage"

	"github.com/spf13/cobra"
)

func main() {

	manager := service.NewTaskService(storage.NewTaskStorage())
	manager.NewDB()

	//migration
	//manager.Migrate()

	// command add
	var addCmd = &cobra.Command{
		Use:   "-a [name] [description]",
		Short: "Add a new task",
		Long:  `Add a new task to the task list with a name and description (Opcional)`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			manager.CreateTask(args[0], args[1])
		},
	}

	// update name and description command
	var updateCmd = &cobra.Command{
		Use:   "-u [id] [name] [description]",
		Short: "Update a task name or description",
		Long: `Update a task name or description with the task id.
		Example:
		# -u 1 "new name" "new description"
		# -u 1 "new name"
		# -u 1 "" "new description"`,
		Args: cobra.RangeArgs(2, 3),
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	// update to in-progress command
	var progressCmd = &cobra.Command{
		Use:   "-p [type] [value]",
		Short: "Update a task to in-progress",
		Long: `Update a task status in-progress with the task id or name.
		type: id or name
		Example:
		# -p id 1
		# -p name "task name
		default value by id"`,
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	// update to done command
	var doneCmd = &cobra.Command{
		Use:   "-d [type] [value]",
		Short: "Update a task to done",
		Long: `Update a task status done with the task id or name.
		type: id or name
		Example:
		# -d id 1
		# -d name "task name"
		default value by id"`,
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	// reset task command
	var resetCmd = &cobra.Command{
		Use:   "-r [type] [value]",
		Short: "Reset a task to todo",
		Long: `Reset a task status to todo with the task id or name.
		type: id or name
		Example:
		# -r id 1
		# -r name "task name"
		default value by id"`,
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	// delete command
	var deleteCmd = &cobra.Command{
		Use:   "-del [type] [value]",
		Short: "Delete a task",
		Long: `Delete a task with the task id or name.
		type: id or name
		Example:
		# -del id 1
		# -del name "task name"
		default value by id"`,
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	// list command
	var listCmd = &cobra.Command{
		Use:   "-l [type] [value]",
		Short: "List tasks",
		Long: `List all tasks, or list task by status, name or id.
		type: id, name, status
		Example:
		# -l
		# -l id 1
		# -l name "task name"
		# -l status "todo"
		default value by satus"`,

		Args: cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	// main command
	var rootCmd = &cobra.Command{
		Use:   "jbot",
		Short: "task tracker CLI",
	}
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(progressCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(resetCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.Execute()

}
