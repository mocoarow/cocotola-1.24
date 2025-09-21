package api_test

import (
	"encoding/json"
	"testing"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/api"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapimb "github.com/mocoarow/cocotola-1.24/lib/api/moonbeam"
)

func TestSpaceID_shouldUnmarshalSuccessfully_whenGivenValidInteger(t *testing.T) {
	t.Parallel()
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
			t.Parallel()
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
	t.Parallel()
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
			t.Parallel()
			var spaceID libapimb.SpaceID
			err := json.Unmarshal([]byte(tt.jsonData), &spaceID)

			if err == nil {
				t.Errorf("SpaceID.UnmarshalJSON() should return error when given %s", tt.name)
			}
		})
	}
}

func TestLang2_shouldUnmarshalSuccessfully_whenGivenValidString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		jsonData string
		expected string
	}{
		{
			name:     "english",
			jsonData: "\"en\"",
			expected: "en",
		},
		{
			name:     "japanese",
			jsonData: "\"ja\"",
			expected: "ja",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var lang2 libapimb.Lang2
			err := json.Unmarshal([]byte(tt.jsonData), &lang2)

			if err != nil {
				t.Errorf("Lang2.UnmarshalJSON() should not return error when given %s, got: %v", tt.name, err)
				return
			}

			if lang2.Value.String() != tt.expected {
				t.Errorf("Lang2.UnmarshalJSON() = %v, want %v", lang2.Value.String(), tt.expected)
			}
		})
	}
}

func TestLang2_shouldReturnError_whenGivenInvalidData(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		jsonData string
	}{
		{
			name:     "invalid json",
			jsonData: "invalid",
		},
		{
			name:     "number",
			jsonData: "123",
		},
		{
			name:     "too long string",
			jsonData: "\"english\"",
		},
		{
			name:     "too short string",
			jsonData: "\"e\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var lang2 libapimb.Lang2
			err := json.Unmarshal([]byte(tt.jsonData), &lang2)

			if err == nil {
				t.Errorf("Lang2.UnmarshalJSON() should return error when given %s", tt.name)
			}
		})
	}
}

func TestAddDeckRequest_shouldUnmarshalSuccessfully_whenGivenValidJSON(t *testing.T) {
	t.Parallel()
	jsonData := `{
		"spaceId": 123,
		"name": "Test Deck",
		"templateId": 456,
		"lang2": "en",
		"description": "Test description"
	}`

	var req api.AddDeckRequest
	err := json.Unmarshal([]byte(jsonData), &req)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if req.SpaceID.Value.Value != 123 {
		t.Errorf("SpaceID = %v, want 123", req.SpaceID.Value.Value)
	}

	if req.Name != "Test Deck" {
		t.Errorf("Name = %v, want 'Test Deck'", req.Name)
	}

	if req.TemplateID.Value.Int() != 456 {
		t.Errorf("TemplateID = %v, want 456", req.TemplateID.Value.Int())
	}

	if req.Lang2.Value.String() != "en" {
		t.Errorf("Lang2 = %v, want 'en'", req.Lang2.Value.String())
	}

	if req.Description != "Test description" {
		t.Errorf("Description = %v, want 'Test description'", req.Description)
	}
}

func TestAddDeckRequest_shouldMarshalSuccessfully_whenGivenValidStruct(t *testing.T) {
	t.Parallel()
	spaceID, err := mbuserdomain.NewSpaceID(123)
	if err != nil {
		t.Fatalf("Failed to create SpaceID: %v", err)
	}

	templateID, err := domain.NewTemplateID(456)
	if err != nil {
		t.Fatalf("Failed to create TemplateID: %v", err)
	}

	lang2, err := mblibdomain.NewLang2("en")
	if err != nil {
		t.Fatalf("Failed to create Lang2: %v", err)
	}

	req := api.AddDeckRequest{
		SpaceID:     libapimb.SpaceID{Value: spaceID},
		Name:        "Test Deck",
		TemplateID:  api.TemplateID{Value: templateID},
		Lang2:       libapimb.Lang2{Value: lang2},
		Description: "Test description",
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	expectedJSON := `{"spaceId":123,"name":"Test Deck","templateId":456,"lang2":"en","description":"Test description"}`
	if string(jsonData) != expectedJSON {
		t.Errorf("Marshal result = %s, want %s", string(jsonData), expectedJSON)
	}
}

func TestAddDeckRequest_shouldReturnError_whenGivenInvalidJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		jsonData string
	}{
		{
			name:     "invalid spaceId type",
			jsonData: `{"spaceId": "invalid", "name": "Test Deck", "templateId": 456, "lang2": "en"}`,
		},
		{
			name:     "invalid templateId type",
			jsonData: `{"spaceId": 123, "name": "Test Deck", "templateId": "invalid", "lang2": "en"}`,
		},
		{
			name:     "invalid lang2 length",
			jsonData: `{"spaceId": 123, "name": "Test Deck", "templateId": 456, "lang2": "english"}`,
		},
		{
			name:     "invalid lang2 type",
			jsonData: `{"spaceId": 123, "name": "Test Deck", "templateId": 456, "lang2": 123}`,
		},
		{
			name:     "malformed json",
			jsonData: `{"spaceId": 123, "name": "Test Deck"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var req api.AddDeckRequest
			err := json.Unmarshal([]byte(tt.jsonData), &req)

			if err == nil {
				t.Errorf("AddDeckRequest.UnmarshalJSON() should return error when %s", tt.name)
			}
		})
	}
}
