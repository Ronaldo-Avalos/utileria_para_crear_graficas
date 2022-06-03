1. Para empezar, necesitamos leer el conjunto de datos(CSV) en nuestro programa. Vamos a crear un nuevo archivo de entrada main.go y luego crear una función principal.

```GO
package main

import (
	"encoding/csv"
	"log"
	"os"
)
func main() {
	// leemos el archivo CSV
	archivocvs, err := os.Open("games.csv")
	var datohorizontal []string
	var datovertical []float64
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(archivocvs)
	reader.LazyQuotes = true
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
}
 ```
#### a. Usamos el ```os``` que es un paquete de Golang para abrir el archivo cvs y el nombre del archivo como argumento sera guardando este conjunto de datos como archivo.
#### b. Declare dos variables para contener los nombres de los juegos y el número de ventas de cada juego( datohorizontal y datovertical).
#### c. Si hay un error al leer el archivo, registraremos el error y saldremos del programa.
