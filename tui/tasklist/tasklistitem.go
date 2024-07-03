package tasklist

import "github.com/burkel24/task-app/pkg/tasks"

type TaskListItem struct {
	task tasks.TaskDTO
}

func (t TaskListItem) Title() string {
	return t.task.Summary
}

func (t TaskListItem) Description() string {
	return ""
}

func (t TaskListItem) FilterValue() string {
	return t.task.Summary
}
