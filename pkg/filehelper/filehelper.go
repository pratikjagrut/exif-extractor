package filehelper

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"os"
)

// CreateFile creates a file with the specified name and format.
func CreateFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// WriteData writes data to the file based on the format.
func WriteData(file *os.File, data [][]string, format string) error {
	switch format {
	case "csv":
		writer := csv.NewWriter(file)
		writer.UseCRLF = true

		err := writer.WriteAll(data)
		if err != nil {
			fmt.Println("Error writing CSV:", err)
			return err
		}

		writer.Flush()

		if err := writer.Error(); err != nil {
			fmt.Println("Error flushing CSV writer:", err)
			return err
		}
	case "html":
		tmpl := template.Must(template.New("table").Parse(`<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<style>
table {
	border-collapse: collapse;
	width: 25%;
}

table, th, td {
	border: 1px solid black;
	padding: 8px;
	text-align: left;
}

</style>
</head>
<body>
<table>
{{range .}}
<tr>{{range .}}
<td>{{.}}</td>{{end}}</tr>
{{end}}
</table>
</body>
</html>`))
		err := tmpl.Execute(file, data)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported file format: %s", format)
	}

	return nil
}
