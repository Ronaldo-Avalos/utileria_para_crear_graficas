#### Comenzamos el proyecto con la creacion del archivo main en Go

```GO
package main

import ()
func main() {
	
}
 ```
 #### Se añadieron las opciones de casos que se presentaran
 
  ```GO
  func main() {
  op1 := os.Args[1]
	op2 := os.Args[2]
	if op2 == "--generate" {
		switch op1 {
		case "--bar":
		case "--line":	
		case "--pie":
		}
	} else {

		switch op1 {
		case "--bar":
		case "--line":
		case "--pie":

		}
	}
}
   ```
#### Generamos una funcion para crear los items randoms para la opcion de grafica de barras
 ```GO
  //Generea la grafica de barras con rand
func generateBarItemsrand() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}
 ```
 ### Generamos la funcion para ordenar los datos de la grafica
  ```GO
  
func barGraph(w http.ResponseWriter, _ *http.Request) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Grafica de barra ",
		Subtitle: "Generada con números aleatorios",
	}))
	bar.SetXAxis([]string{"dato1", "dato2", "dato3", "dato4", "dato5", "dato6", "dato7"}).
		AddSeries("Category A", generateBarItemsrand()).
		AddSeries("Category B", generateBarItemsrand())
	bar.Render(w)
}
 ```
  
