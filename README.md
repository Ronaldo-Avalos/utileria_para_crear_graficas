#### Para empezar, necesitamos leer el conjunto de datos(CSV) en nuestro programa. Vamos a crear un nuevo archivo de entrada main.go y luego crear una función principal.

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

#### Cabe señalar aquí que el conjunto de datos CSV se leerá como  ```recordsy ``` cada fila se representará como un elemento de matriz, es decir, el encabezado será el primer elemento de la matriz y, posteriormente, cada fila seguirá su ejemplo.

#### Lo siguiente es formatear los registros obtenidos. Creamos una nueva función. La función toma recordwhich is a slice of string slice  ```[][]stringy ``` devuelve dos variables con el mismo tipo de dato
