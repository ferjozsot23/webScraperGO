package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func SaveDataOnCSVFormat(data [][]string, filename string) {
	file, err := os.Create(fmt.Sprintf("%s.csv", filename))

	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()
	csvWriter := csv.NewWriter(file)

	for _, rowData := range data {
		csvWriter.Write(rowData)
	}
	// Lo de Buffer finalmente se escribe en "file.csv"
	csvWriter.Flush()

}
