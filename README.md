# EXIF Extractor

EXIF Extractor is a command-line utility written in Go that reads the EXIF data from images in a specified directory and generates a CSV or HTML file containing the image file path, GPS position latitude, and GPS position longitude.

## Installation

Clone the repository:

```
git clone https://github.com/your-username/exif-extractor.git
```

Change to the project directory:

```
cd exif-extractor
```

Build the executable:

```
go build
```

## Usage

```
./exif-extractor [flags]
```

The following flags are available:

- `--dir`: Specify the directory to process. Default is the current directory.
- `--out`: Specify the output file path. Can be 'csv' or 'html'. Default is 'result.csv'.
- `--logs`: Set false if you want to disable logs. Default is 'true'.

### Examples

Process images in the current directory and generate a CSV file:

```
./exif-extractor
```

Process images in a specific directory and generate an HTML file:

```
./exif-extractor --dir=/path/to/images --out=result.html
```

## Features

- Reads EXIF data from images in the specified directory
- Supports JPEG/JPG image formats
- Writes the image file path, GPS position latitude, and GPS position longitude to a CSV or HTML file
- Logs progress and errors during the process

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.