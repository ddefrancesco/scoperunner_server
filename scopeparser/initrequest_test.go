package scopeparser

import (
	"strings"
	"testing"
	"time"
)

func TestInitializeRequest_SetTimeCommand(t *testing.T) {
	s := &InitializeRequest{}
	cmd := s.SetTimeCommand()

	// Test that the command starts with the expected prefix
	if !strings.HasPrefix(cmd, ":SL") {
		t.Errorf("SetTimeCommand() returned an unexpected command: %s", cmd)
	}

	// Test that the command has the expected length
	expectedLength := len(":SL") + len("15:04:05#")
	if len(cmd) != expectedLength {
		t.Errorf("SetTimeCommand() returned a command with unexpected length: %d, expected %d", len(cmd), expectedLength)
	}

	// Test that the time portion of the command is in the expected format
	timePortion := cmd[len(":SL"):]
	_, err := time.Parse("15:04:05#", timePortion)
	if err != nil {
		t.Errorf("SetTimeCommand() returned an invalid time portion: %s", timePortion)
	}

	// Test that the time portion is in the expected time zone
	loc, _ := time.LoadLocation("Europe/Rome")
	now := time.Now().In(loc)
	expectedTime := now.Format("15:04:05#")
	if timePortion != expectedTime {
		t.Errorf("SetTimeCommand() returned an unexpected time portion: %s, expected %s", timePortion, expectedTime)
	}
}
