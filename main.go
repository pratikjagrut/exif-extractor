package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/pratikjagrut/exif-extractor/pkg/dms"
	"github.com/pratikjagrut/exif-extractor/pkg/exif"
	"github.com/pratikjagrut/exif-extractor/pkg/filehelper"
	"github.com/pratikjagrut/exif-extractor/pkg/log"
)

type config struct {
	directory  string
	output     string
	enableLogs bool
}

var paths []string

func parseFlags() *config {
	cfg := &config{}

	flag.StringVar(&cfg.directory, "dir", ".", "Specify the directory to process. Default is current directory.")
	flag.StringVar(&cfg.output, "out", "result.csv", "Specify the output file path. Can be 'csv' or 'html'. Default is 'result.csv'.")
	flag.BoolVar(&cfg.enableLogs, "logs", true, "Set false if you want to disable logs. Default is 'true'")

	// Override the default help print
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("    %s --directory=/path/to/images --out=result.html\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	return cfg
}

func visit(path string, d fs.DirEntry, err error) error {
	if d.IsDir() {
		return nil
	} else {
		ext := strings.Split(path, ".")[1]
		if ext != "jpg" && ext != "jpeg" {
			log.Warning("Unsupported image format: %s", path)
		} else {
			paths = append(paths, path)
		}
	}
	return nil
}

func main() {
	cfg := parseFlags()
	outFile := strings.Split(cfg.output, ".")
	if len(outFile) == 0 || len(outFile) > 2 {
		log.Fatal("Invalid output file format")
	}
	format := strings.ToLower(outFile[1])
	if format != "csv" && format != "html" {
		log.Fatal(`Output file format should be either "csv" or "html"`)
	}

	log.EnableLogs = cfg.enableLogs

	log.Info("Directory to process: %s", cfg.directory)
	log.Info("Output file location: %s", cfg.output)

	// Check if the directory exists
	_, err := os.Stat(cfg.directory)
	if os.IsNotExist(err) {
		log.Fatal("Directory not found: %s", cfg.directory)
	}

	// Traverse directory, read EXIF data, and write to output here.
	err = filepath.WalkDir(cfg.directory, visit)
	if err != nil {
		log.Fatal("Error walking the path %s: %v", cfg.directory, err)
	}

	data := make([][]string, 0)
	data = append(data, []string{"Path", "Latitude", "Longitude"})

	for _, path := range paths {
		log.Info("Processing image: %s", path)
		latitude, longitude, err := exif.GetLatLong(path)
		if err != nil {
			log.Error("Failed to get latitude and longitude for image: %s", path)
			continue
		}
		gpsLatitude, gpsLongitude, err := dms.NewDMS(latitude, longitude)
		if err != nil {
			log.Error("Failed to convert latitude and longitude to GPS coordinates: %v", err)
			continue
		}
		data = append(data, []string{path, gpsLatitude.String(), gpsLongitude.String()})
	}

	// Create the file
	file, err := filehelper.CreateFile(cfg.output)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	err = filehelper.WriteData(file, data, format)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info("Successfully created file with EXIF information at location: %s", cfg.output)
}
