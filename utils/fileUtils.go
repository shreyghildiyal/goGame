package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	config "github.com/shreyghildiyal/goGame/configs"
)

func GetCsvRecords(filePath string) ([][]string, error) {

	f, err := os.Open(filePath)
	if err != nil {
		// log.Fatal("Unable to read input file "+filePath, err)
		log.Printf("Unable to open input file "+filePath, err)
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Printf("Unable to parse file as CSV for "+filePath, err)
		return nil, err
	}
	return records, nil
}

func SaveToFile(jsonStr []byte, fileName string) error {

	saveDir := config.GetConfig().SaveGameDir
	filePath := fmt.Sprintf("%s/%s", saveDir, fileName)

	err := os.WriteFile(filePath, jsonStr, 0644)

	return err

}
