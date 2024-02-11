package main

import (
	"fmt"
	"github.com/robertszeker/snake-on-max-7219-go/matrix"
	"github.com/robertszeker/snake-on-max-7219-go/object"
	"log/slog"
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
		return fmt.Errorf("failed to initialize board: %w", err)
	}
	defer m.Close()

	mouse := object.NewMouse(object.NewPoint(1, 2))
	m.AddObject(mouse)
	err = m.PrintObjects()
	if err != nil {
		return fmt.Errorf("failed to write point: %w", err)
	}

	return nil
}
