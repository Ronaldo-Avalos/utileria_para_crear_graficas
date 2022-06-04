package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	matriz := []string{}
	file, err := os.Open("./data.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for filas, record := range records {
		for _, item := range record {
			matriz[filas] = item
		}
		fmt.Println(matriz)
	}
}
