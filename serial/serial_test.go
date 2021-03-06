package serial_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/esp32-controller/serial"
	"testing"
)

func TestSerialBegin(t *testing.T) {
	baud := serial.BaudRate115200
	serial.Baud = 0
	serial.Begin(baud)
	assert.That(t, serial.Baud, is.Equal(serial.BaudRate115200))
}
