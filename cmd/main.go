package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/burkel24/task-app/pkg/db"
)

func main() {
	_, err := db.InitDb()

	if err != nil {
		slog.Error("DB connection failed %w", slog.Attr{Key: "error", Value: slog.AnyValue(err)})
		os.Exit(1)
	}

	fmt.Printf("Hello2")
}
