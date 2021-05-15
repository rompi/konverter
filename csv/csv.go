package main

import (
	//"bufio"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func openFile(filePath string) (*os.File, error) {
	// Open the file
	return os.Open(filePath)
}

func parseCsv(csvFile *os.File) ([][]string, error) {
	// Parse the file
	csvReader := csv.NewReader(csvFile)
	return csvReader.ReadAll()
}

func main() {
	csvFile, err := openFile("files/input1.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer csvFile.Close()
	records, err := parseCsv(csvFile)
	if err != nil {
		panic(err)
	}

	type player struct {
		No       string `json:"no"`
		Name     string `json:"name"`
		Position string `json:"position"`
	}

	var players []player
	for i, record := range records {
		if i == 0 {
			continue
		}
		players = append(players, player{
			No:       record[0],
			Name:     record[1],
			Position: record[2],
		})
	}

	file, _ := json.MarshalIndent(players, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)
}
