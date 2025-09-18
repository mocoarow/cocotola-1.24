package controller_test

import (
	"encoding/json"
	"testing"

	libapimb "github.com/mocoarow/cocotola-1.24/lib/api/moonbeam"
)

func TestSpaceID_shouldUnmarshalSuccessfully_whenGivenValidInteger(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		expected int
	}{
		{
			name:     "positive integer",
			jsonData: "123",
			expected: 123,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var spaceID libapimb.SpaceID
			err := json.Unmarshal([]byte(tt.jsonData), &spaceID)

			if err != nil {
				t.Errorf("SpaceID.UnmarshalJSON() should not return error when given %s, got: %v", tt.name, err)
				return
			}

			if spaceID.Value.Value != tt.expected {
				t.Errorf("SpaceID.UnmarshalJSON() = %v, want %v", spaceID.Value.Value, tt.expected)
			}
		})
	}
}

func TestSpaceID_shouldReturnError_whenGivenInvalidData(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
	}{
		{
			name:     "invalid json",
			jsonData: "invalid",
		},
		{
			name:     "string number",
			jsonData: "\"123\"",
		},
		{
			name:     "zero value",
			jsonData: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var spaceID libapimb.SpaceID
			err := json.Unmarshal([]byte(tt.jsonData), &spaceID)

			if err == nil {
				t.Errorf("SpaceID.UnmarshalJSON() should return error when given %s", tt.name)
			}
		})
	}
}

func TestSpaceIDs_shouldUnmarshalSuccessfully_whenGivenValidArray(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		expected []int
	}{
		{
			name:     "single id",
			jsonData: "[123]",
			expected: []int{123},
		},
		{
			name:     "multiple ids",
			jsonData: "[123, 456, 789]",
			expected: []int{123, 456, 789},
		},
		{
			name:     "empty array",
			jsonData: "[]",
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var spaceIDs libapimb.SpaceIDs
			err := json.Unmarshal([]byte(tt.jsonData), &spaceIDs)

			if err != nil {
				t.Errorf("SpaceIDs.UnmarshalJSON() should not return error when given %s, got: %v", tt.name, err)
				return
			}

			if len(spaceIDs) != len(tt.expected) {
				t.Errorf("SpaceIDs length = %v, want %v", len(spaceIDs), len(tt.expected))
				return
			}

			for i, expected := range tt.expected {
				if spaceIDs[i] != expected {
					t.Errorf("SpaceIDs[%d] = %v, want %v", i, spaceIDs[i], expected)
				}
			}
		})
	}
}

func TestSpaceIDs_shouldReturnError_whenGivenInvalidData(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
	}{
		{
			name:     "invalid json",
			jsonData: "invalid",
		},
		{
			name:     "string array",
			jsonData: `["123", "456"]`,
		},
		{
			name:     "mixed types",
			jsonData: `[123, "456"]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var spaceIDs libapimb.SpaceIDs
			err := json.Unmarshal([]byte(tt.jsonData), &spaceIDs)

			if err == nil {
				t.Errorf("SpaceIDs.UnmarshalJSON() should return error when given %s", tt.name)
			}
		})
	}
}

func TestSpaceIDs_ToSpaceIDs_shouldReturnValidSpaceIDs(t *testing.T) {
	spaceIDs := libapimb.SpaceIDs{123, 456}

	result, err := spaceIDs.ToSpaceIDs()
	if err != nil {
		t.Fatalf("ToSpaceIDs() should not return error, got: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("ToSpaceIDs() length = %v, want 2", len(result))
		return
	}

	if result[0].Value != 123 {
		t.Errorf("result[0].Value = %v, want 123", result[0].Value)
	}

	if result[1].Value != 456 {
		t.Errorf("result[1].Value = %v, want 456", result[1].Value)
	}
}