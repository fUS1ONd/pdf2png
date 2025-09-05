package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

// Config содержит параметры для операции конвертации.
type Config struct {
	InputPath string
	OutputDir string
	DPI       float64
	Verbose   bool
}

// Convert выполняет основную логику конвертации PDF в PNG.
func Convert(cfg Config) error {
	doc, err := fitz.New(cfg.InputPath)
	if err != nil {
		return fmt.Errorf("failed to open PDF: %w", err)
	}
	defer doc.Close()

	if err := os.MkdirAll(cfg.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	if cfg.Verbose {
		fmt.Printf("Processing PDF: %s\n", cfg.InputPath)
		fmt.Printf("Output directory: %s\n", cfg.OutputDir)
		fmt.Printf("DPI: %.0f\n", cfg.DPI)
		fmt.Printf("Total pages: %d\n", doc.NumPage())
	}

	for n := 0; n < doc.NumPage(); n++ {
		pageNum := n + 1
		if cfg.Verbose {
			fmt.Printf("Processing page %d/%d...\n", pageNum, doc.NumPage())
		}

		img, err := doc.Image(n)
		if err != nil {
			return fmt.Errorf("failed to render page %d: %w", pageNum, err)
		}

		scaledImg := scaleImage(img, cfg.DPI)

		outputPath := filepath.Join(cfg.OutputDir, fmt.Sprintf("page_%03d.png", pageNum))

		if err := saveImage(scaledImg, outputPath); err != nil {
			return fmt.Errorf("failed to save page %d: %w", pageNum, err)
		}

		if cfg.Verbose {
			fmt.Printf("  → Saved: %s\n", outputPath)
		}
	}

	return nil
}

func scaleImage(img image.Image, dpi float64) image.Image {
	const baseDPI = 150.0
	scale := dpi / baseDPI

	if scale == 1.0 {
		return img
	}

	bounds := img.Bounds()
	width := int(float64(bounds.Dx()) * scale)
	height := int(float64(bounds.Dy()) * scale)

	scaled := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			srcX := int(float64(x) / scale)
			srcY := int(float64(y) / scale)
			scaled.Set(x, y, img.At(srcX, srcY))
		}
	}

	return scaled
}

func saveImage(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
