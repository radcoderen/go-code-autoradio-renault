package codeautoradio

import (
	"testing"
)

// precodeIsValid checks if the given precode is valid.
func TestPrecodeIsValid(t *testing.T) {
	tests := []struct {
		precode string
		want    bool
	}{
		{"A123", true},
		{"B456", true},
		{"C789", true},
		{"Z999", true},
		{"A1234", false},
		{"A12", false},
		{"A012", false},
		{"0123", false},
	}
	for _, tt := range tests {
		t.Run(tt.precode, func(t *testing.T) {
			if got := precodeIsValid(tt.precode); got != tt.want {
				t.Errorf("precodeIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// GetRadioCode returns the radio code for the given precode.
func TestGetRadioCode(t *testing.T) {
	tests := []struct {
		precode string
		want    string
		ok      bool
	}{
		{"A123", "0086", true},
		{"B456", "0709", true},
		{"C789", "1621", true},
		{"Z999", "0060", true},
		{"A1234", "", false},
		{"A12", "", false},
		{"A012", "", false},
		{"0123", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.precode, func(t *testing.T) {
			got, ok := GetRadioCode(tt.precode)
			if got != tt.want {
				t.Errorf("GetRadioCode() got = %v, want %v", got, tt.want)
			}
			if ok != tt.ok {
				t.Errorf("GetRadioCode() ok = %v, want %v", ok, tt.ok)
			}
		})
	}
}
