package main

import (
	"github.com/adrianh-za/go-max7219"
	"log"
	"log/slog"
	"time"
)

func main() {
	if err := mainInternal(); err != nil {
		slog.Error(err.Error())
	} else {
		slog.Info("DONE")
	}
}

func mainInternal() error {
	// Create new LED matrix with number of cascaded devices is equal to 1
	mtx := max7219.NewMatrix(4, max7219.RotateAntiClockwiseInvert)
	// Open SPI device with spibus and spidev parameters equal to 0 and 0.
	// Set LED matrix brightness is equal to 7
	err := mtx.Open(0, 0, 7)
	if err != nil {
		log.Fatal(err)
	}
	defer mtx.Close()
	// Output text message to LED matrix
	mtx.SlideMessage("Hello world!!! Hey Beavis, let's rock!",
		max7219.FontCP437, true, 50*time.Millisecond)
	// Wait 1 sec, then continue output new text
	time.Sleep(1 * time.Second)
	// Output national text (russian in example) to LED matrix
	mtx.SlideMessage("Привет мир!!! Шарик, ты - балбес!!!",
		max7219.FontZXSpectrumRus, true, 50*time.Millisecond)
	return nil
}
