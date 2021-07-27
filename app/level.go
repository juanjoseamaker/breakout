package app

import (
	"encoding/csv"
	"os"
	"log"
	"strconv"
)

func LoadLevel(arr [][]uint8, filename string) {
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i := range csvLines {
		for j := range csvLines[i] {
			num, err := strconv.ParseInt(csvLines[i][j], 10, 8)
			arr[i][j] = uint8(num)
			
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
