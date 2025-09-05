package main

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestScaleImage(t *testing.T) {
	originalImg := image.NewRGBA(image.Rect(0, 0, 10, 20))

	dpi := 300.0

	scaledImg := scaleImage(originalImg, dpi)

	expectedWidth := 20
	expectedHeight := 40

	bounds := scaledImg.Bounds()
	if bounds.Dx() != expectedWidth || bounds.Dy() != expectedHeight {
		t.Errorf("expected dimensions %dx%d, but got %dx%d",
			expectedWidth, expectedHeight, bounds.Dx(), bounds.Dy())
	}
}

func TestConvert(t *testing.T) {
	tempDir := t.TempDir()

	config := Config{
		InputPath: filepath.Join("testdata", "sample.pdf"),
		OutputDir: tempDir,
		DPI:       150,
		Verbose:   false,
	}

	if _, err := os.Stat(config.InputPath); os.IsNotExist(err) {
		t.Fatalf("Test PDF file not found at %s. Please create it.", config.InputPath)
	}

	err := Convert(config)
	if err != nil {
		t.Fatalf("Convert() failed with error: %v", err)
	}

	expectedFileName := "page_001.png"
	expectedFilePath := filepath.Join(tempDir, expectedFileName)

	if _, err := os.Stat(expectedFilePath); os.IsNotExist(err) {
		t.Fatalf("Expected output file was not created: %s", expectedFilePath)
	}

	file, err := os.Open(expectedFilePath)
	if err != nil {
		t.Fatalf("Could not open created file: %v", err)
	}
	defer file.Close()

	_, err = png.Decode(file)
	if err != nil {
		t.Fatalf("Created file is not a valid PNG: %v", err)
	}

	dirEntries, err := os.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("Could not read output directory: %v", err)
	}
	if len(dirEntries) != 1 {
		t.Errorf("Expected 1 file in output directory, but found %d", len(dirEntries))
	}
}
