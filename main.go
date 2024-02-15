package main

import (
	"fmt"
	"github.com/robertszeker/snake-on-max-7219-go/matrix"
	"github.com/robertszeker/snake-on-max-7219-go/object"
	"log/slog"
	"time"
)

// todo: create Makefile to build, send and run remotely

func main() {
	if err := mainInternal(); err != nil {
		slog.Error(err.Error())
	} else {
		slog.Info("DONE")
	}
}

func mainInternal() error {

	m, err := matrix.NewMatrix(4)
	if err != nil {
		return fmt.Errorf("failed to initialize matrix: %w", err)
	}
	defer m.Close()

	err = m.PrintObjects()

	snake := object.NewSnake(object.NewPoint(1, 2), m.PrintObjects)
	m.AddObject(snake)

	time.Sleep(10 * time.Second)

	if err != nil {
		return fmt.Errorf("failed to write point: %w", err)
	}

	return nil
}
