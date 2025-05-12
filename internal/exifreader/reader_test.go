package exifreader

import "testing"

func TestExtractCameraModel(t *testing.T) {
	testFile := "../../testdata/Nikon_D70.jpg"
	expectedModel := "NIKON D70"

	model, err := ExtractCameraModel(testFile)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if model != expectedModel {
		t.Errorf("Expected %s, got %s", expectedModel, model)
	}
}
