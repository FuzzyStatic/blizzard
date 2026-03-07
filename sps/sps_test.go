package sps

import (
	"encoding/json"
	"testing"
)

func TestPlayableTitles_Unmarshal(t *testing.T) {
	jsonData := `{
	    "titles": [
	        {"id": 1095849281, "name": "Avowed"},
	        {"id": 1381257807, "name": "Blizzard Arcade Collection"}
	    ]
	}`

	var pt PlayableTitles
	err := json.Unmarshal([]byte(jsonData), &pt)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if len(pt.Titles) != 2 {
		t.Errorf("Expected 2 titles, got %d", len(pt.Titles))
	}

	if pt.Titles[0].ID != 1095849281 {
		t.Errorf("Expected first title ID 1095849281, got %d", pt.Titles[0].ID)
	}
}
