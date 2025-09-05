package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		inputPath string
		outputDir string
		dpi       float64
		verbose   bool
	)

	flag.StringVar(&inputPath, "input", "", "Path to input PDF file (required)")
	flag.StringVar(&inputPath, "i", "", "Path to input PDF file (shorthand)")
	flag.StringVar(&outputDir, "output", "", "Output directory (optional, default: PDF filename without extension)")
	flag.StringVar(&outputDir, "o", "", "Output directory (shorthand)")
	flag.Float64Var(&dpi, "dpi", 150, "DPI for output images")
	flag.BoolVar(&verbose, "verbose", false, "Verbose output")
	flag.BoolVar(&verbose, "v", false, "Verbose output (shorthand)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "PDF to PNG Converter\n")
		fmt.Fprintf(os.Stderr, "Usage: %s -i <input.pdf> [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  %s -i document.pdf -o images -dpi 300\n", os.Args[0])
	}

	flag.Parse()

	if inputPath == "" {
		fmt.Fprintf(os.Stderr, "Error: input PDF file is required\n\n")
		flag.Usage()
		os.Exit(1)
	}
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: input file '%s' does not exist\n", inputPath)
		os.Exit(1)
	}

	if outputDir == "" {
		base := filepath.Base(inputPath)
		outputDir = strings.TrimSuffix(base, filepath.Ext(base))
	}

	config := Config{
		InputPath: inputPath,
		OutputDir: outputDir,
		DPI:       dpi,
		Verbose:   verbose,
	}

	if err := Convert(config); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully converted '%s' to PNG images in '%s'\n", inputPath, outputDir)
}
