package converter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func ConvertJSONToCSV(inputPath, outputPath string) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	var jsonArray []map[string]interface{}
	if err := json.Unmarshal(data, &jsonArray); err != nil {
		return fmt.Errorf("unmarshaling JSON: %w", err)
	}

	if len(jsonArray) == 0 {
		return fmt.Errorf("no data found in JSON")
	}

	var headers []string
	for k := range jsonArray[0] {
		headers = append(headers, k)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("creating CSV file: %w", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("writing headers: %w", err)
	}

	for _, obj := range jsonArray {
		row := make([]string, len(headers))
		for i, header := range headers {
			row[i] = fmt.Sprintf("%v", obj[header])
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("writing row: %w", err)
		}
	}

	return nil
}
