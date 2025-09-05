# pdf2png

`pdf2png` is a fast, cross-platform command-line tool written in Go that converts PDF documents into high-quality PNG images. It provides adjustable DPI settings for full quality control.

## Features

* Converts each page of a PDF into a separate PNG image.
* Adjustable DPI for controlling output image quality.
* Cross-platform support (Windows, macOS, Linux).
* Written in Go for simplicity and performance.
* Includes tests to ensure reliability.

## Installation

### Using Go

If you have Go installed, you can build and install `pdf2png` directly:

```bash
go install github.com/fUS1ONd/pdf2png@latest
```

This command will download the source code, build the application, and place the executable in your Go bin directory.

### Using Git Clone

You can also clone the repository and build it manually:

```bash
git clone https://github.com/fUS1ONd/pdf2png.git
cd pdf2png
go build -o pdf2png
```

After building, the `pdf2png` executable will be available in the current directory.

### Using Precompiled Binaries

Precompiled binaries for various platforms are available in the [releases section](https://github.com/fUS1ONd/pdf2png/releases). Download the appropriate binary for your operating system and architecture.

## Usage

After installation, you can use `pdf2png` from the command line:

```bash
pdf2png input.pdf output_directory
```

This command will convert each page of `input.pdf` into a separate PNG image and save them in the specified `output_directory`.

### Options

* `-dpi`: Set the DPI (dots per inch) for the output images. Higher DPI values result in higher quality images.

```bash
pdf2png -dpi 300 input.pdf output_directory
```

* `-help`: Display help information.

```bash
pdf2png -help
```

## Testing

To run tests:

```bash
go test ./...
```

This will execute all available tests to ensure the tool works as expected.

## License

This project is licensed under the GNU General Public License v2.0.

## Contributing

Contributions are welcome! Please fork the repository, make your changes, and submit a pull request.

## Contact

For issues or questions, please open an issue on the [GitHub repository](https://github.com/fUS1ONd/pdf2png/issues).
