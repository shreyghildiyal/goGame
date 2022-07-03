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
	var filePath string
	if len(saveDir) == 0 {
		filePath = fileName
	} else {
		filePath = fmt.Sprintf("%s/%s", saveDir, fileName)
	}

	err := createDirectoryStructure(saveDir)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err0 := os.Create(filePath)
	if err0 != nil {
		fmt.Printf("Error creating path %s \n", filePath)
		fmt.Println(err0)
		return fmt.Errorf("error creating the file %s", filePath)
	}
	err = os.WriteFile(filePath, jsonStr, 0644)
	if err != nil {
		fmt.Printf("Error writing to path %s \n", filePath)
	} else {
		fmt.Printf("The file should be written to %s\n", filePath)
	}

	return err

}

func createDirectoryStructure(dir string) error {

	err := os.MkdirAll(dir, os.ModePerm)
	return err
}
