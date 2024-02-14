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

	mouse1 := object.NewMouse(object.NewPoint(1, 2))
	mouse2 := object.NewMouse(object.NewPoint(3, 2))

	m.AddObject(mouse1)
	m.AddObject(mouse2)

	err = m.PrintObjects()

	time.Sleep(2 * time.Second)

	slog.Info("moving mouse ...")
	mouse1.Change(object.NewPoint(4, 6))
	err = m.PrintObjects()

	snake := object.NewSnake(object.NewPoint(1, 2), m.GetPrintListener())
	m.AddObject(snake)

	if err != nil {
		return fmt.Errorf("failed to write point: %w", err)
	}

	return nil
}
