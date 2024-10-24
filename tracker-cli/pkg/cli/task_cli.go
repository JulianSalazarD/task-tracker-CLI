package cli

import "tracker-cli/internal/service"

type CLI struct {
	taskService *service.TaskService
}

func NewCLI(taskService *service.TaskService) *CLI {
	return &CLI{taskService: taskService}
}
