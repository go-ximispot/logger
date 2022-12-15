package logger

import (
	"bytes"
	"testing"
)

func TestDebug(t *testing.T) {
	// Create a bytes.Buffer to use as the logger's output
	buf := &bytes.Buffer{}

	// Create a new Logger with DebugLevel log level
	logger := NewLogger(buf, DebugLevel)

	// Call the Debug method with a test message
	logger.Debug("this is a test message")

	// Check that the correct log message was written to the output
	if buf.String() != "DEBUG: this is a test message\n" {
		t.Errorf("unexpected log message: %s", buf.String())
	}
}
