Para empezar, necesitamos leer el conjunto de datos(CSV) en nuestro programa. Vamos a crear un nuevo archivo de entrada main.go y luego crear una función principal.

```GO
package main

import (
	"encoding/csv"
	"log"
	"os"
)
func main() {
	// CSV
	file, err := os.Open("games.csv")
	var gameNames []string
	var sales []float64
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
}
 ```
Usamos el ```os``` que es un paquete de Golang para abrir el archivo cvs y el nombre del archivo como argumento sera guardando este conjunto de datos como archivo
