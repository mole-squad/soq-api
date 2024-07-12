package common

type AppState int

const (
	AppStateLogin AppState = iota
	AppStateTaskList
	AppStateTaskForm
)
