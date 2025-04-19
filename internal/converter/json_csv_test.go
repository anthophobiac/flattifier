package converter

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"testing"
)

func TestConvertJSONToCSV(t *testing.T) {
	tmpDir := t.TempDir()

	inputFile := filepath.Join(tmpDir, "input.json")
	outputFile := filepath.Join(tmpDir, "output.csv")

	jsonData := `[
		{"name": "Johnnie", "age": 30, "email": "johnnie@human.com"},
		{"name": "Archie", "age": 8, "email": "archie@cat.com"}
	]`

	if err := os.WriteFile(inputFile, []byte(jsonData), 0644); err != nil {
		t.Fatalf("Failed to write input file: %v", err)
	}

	err := ConvertJSONToCSV(inputFile, outputFile)
	if err != nil {
		t.Fatalf("ConvertJSONToCSV failed: %v", err)
	}

	file, err := os.Open(outputFile)
	if err != nil {
		t.Fatalf("Failed to open output CSV: %v", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read CSV: %v", err)
	}

	if len(records) != 3 {
		t.Fatalf("Expected 3 rows (including header), got %d", len(records))
	}

	expectedHeaders := []string{"name", "age", "email"}
	for i, header := range expectedHeaders {
		if records[0][i] != header {
			t.Errorf("Expected header %q, got %q", header, records[0][i])
		}
	}

	expectedFirstRow := []string{"Johnnie", "30", "johnnie@human.com"}
	for i, val := range expectedFirstRow {
		if records[1][i] != val {
			t.Errorf("Expected row value %q, got %q", val, records[1][i])
		}
	}
}
