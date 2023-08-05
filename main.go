package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// ReadCsv reads a CSV file and returns a slice of its lines.
func ReadCsv(folderPath string) (map[string][]string, error) {
	// Open the folder
	folder, err := os.Open(folderPath)
	if err != nil {
		log.Fatal("Unable to read and open Folder path "+folderPath, err)
		return nil, err
	}
	defer folder.Close()

	// Read the names of files in the folder
	filenames, err := folder.Readdirnames(-1)
	if err != nil {
		log.Fatal("Unable to read folder contents: ", err)
		return nil, err
	}

	data := make(map[string][]string)

	for _, filename := range filenames {
		if filepath.Ext(filename) == ".csv" {
			filepath := filepath.Join(folderPath, filename)
			fmt.Println("Reading", filepath)

			// Open the CSV file
			file, err := os.Open(filepath)
			if err != nil {
				log.Println("Unable to read and open file "+filepath, err)
				continue
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)

			fmt.Println(scanner)

			for scanner.Scan() {
				line := scanner.Text()
				// Split the line into fields
				currentLine := strings.Split(strings.Split(line, "\n")[0], ",")				
				col_id := currentLine[0]
				if existingValues, exist := data[col_id]; exist {
					mergedValues := append(existingValues, currentLine[1:]...)
					data[col_id] = mergedValues
				} else {
					data[col_id] = currentLine[1:]
				}
			}

			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading file:", err)
			}
		}
	}
	return data, nil
}

func WriteCsv(data map[string][]string, filepath string) error {
	
	csvFile, err := os.Create(filepath)

	if err != nil {
		log.Fatal("Unable to read and open Folder path "+filepath, err)
		return err
	}

	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	for key, value := range data {

		line := append([]string{key}, value...)
		err := writer.Write(line)

		if err != nil {
			log.Fatal("Unable to write to file "+filepath, err)
			return err
		}
	}

	
	return nil

}

func main() {
	records, err := ReadCsv("C:\\Users\\amsam\\go-test\\csv_dev")
	fmt.Println(err)
	WriteCsv(records, "C:\\Users\\amsam\\go-test\\csv_dev\\merged.csv")
}