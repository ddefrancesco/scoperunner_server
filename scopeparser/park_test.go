package scopeparser

import (
	"testing"
)

func TestParkTypeConstants(t *testing.T) {
	tests := []struct {
		name     string
		parkType ParkType
		expected string
	}{
		{
			name:     "ParkTypeQuery",
			parkType: ParkTypeQuery,
			expected: "query",
		},
		{
			name:     "ParkTypeSeek",
			parkType: ParkTypeSeek,
			expected: "seek",
		},
		{
			name:     "ParkTypePark",
			parkType: ParkTypePark,
			expected: "park",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.parkType) != tt.expected {
				t.Errorf("ParkType %s = %v, want %v", tt.name, tt.parkType, tt.expected)
			}
		})
	}
}

func TestParkTypeUniqueness(t *testing.T) {
	parkTypes := []ParkType{ParkTypeQuery, ParkTypeSeek, ParkTypePark}
	seen := make(map[ParkType]bool)

	for _, pt := range parkTypes {
		if seen[pt] {
			t.Errorf("Duplicate ParkType found: %v", pt)
		}
		seen[pt] = true
	}
}

func TestParkTypeValidity(t *testing.T) {
	invalidParkType := ParkType("invalid")
	validParkTypes := map[ParkType]bool{
		ParkTypeQuery: true,
		ParkTypeSeek:  true,
		ParkTypePark:  true,
	}

	if validParkTypes[invalidParkType] {
		t.Errorf("Invalid ParkType %v should not be recognized as valid", invalidParkType)
	}
}
