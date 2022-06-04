package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Archivo de texto con encabezado

func main() {
	file, err := os.Create("data.csv")
	if err != nil {
		fmt.Println(err)
	}
	writer := csv.NewWriter(file)
	var data = [][]string{
		{"Enero", "Febrero", "Marzo", "Abril", "Junio", "Julio"},
		{"23", "17", "43", "23", "23", "12"},
		{"12", "16", "22", "34", "12", "33"},
		{"5", "23", "24", "33", "11", "22"},
		{"34", "14", "25", "31", "23", "12"},
		{"22", "14", "26", "32", "45", "23"},
	}
	err = writer.WriteAll(data)
	if err != nil {
		fmt.Println(err)
	}
}
