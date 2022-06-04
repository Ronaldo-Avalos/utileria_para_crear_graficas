package main

import (
	"encoding/csv"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

//Generea la grafica de barras con rand
func generateBarItemsrand() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

//Generea la grafica de barras con rand
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

//Genera la grafica de lineas con rand
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

//Genera la grafica de lineas con rand
func httpserver(w http.ResponseWriter, _ *http.Request) {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Grafica de lineas",
			Subtitle: "Generada con numeros aleatorios",
		}))
	line.SetXAxis([]string{"dato1", "dato2", "dato3", "dato4", "dato5", "dato6", "dato7"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

//Genera la grafica de pasatel con rand
func pieGraph(w http.ResponseWriter, _ *http.Request) {
	destinations := []opts.PieData{
		{Name: "dato1", Value: rand.Intn(300)},
		{Name: "dato2", Value: rand.Intn(300)},
		{Name: "dato2", Value: rand.Intn(300)},
		{Name: "dato3", Value: rand.Intn(300)},
		{Name: "dato4", Value: rand.Intn(300)},
	}
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeChalk}),
		charts.WithTitleOpts(opts.Title{Title: "Grafica de pastele"}),
	)
	pie.AddSeries("destinations", destinations)
	pie.Render(w)
}

//Leemos el archivo CSV para mandarlo a componer en una matriz

func ReadCSVFile() [][]string {
	// CSV Reader
	file, err := os.Open("./data.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

// Generamos la matriz con los datos del CSV
func generateBarFile() []opts.BarData {
	items := make([]opts.BarData, 0)
	sum := 0

	for i := 0; i < len(ReadCSVFile()[0]); i++ {
		sum = 0
		for j := 1; j < len(ReadCSVFile()); j++ {
			temp, _ := strconv.Atoi(ReadCSVFile()[j][i])
			sum += temp
		}
		value := strconv.Itoa(sum)
		items = append(items, opts.BarData{Value: value})
	}
	return items
}

// Generamos la grafica de barras con el archivo CSV
func barGraphcsv(w http.ResponseWriter, _ *http.Request) {

	bar := charts.NewBar()
	//opciones globales
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "GRAFICA DE BARRAS",
		Subtitle: "Grafica generada con archivo csv",
	}))

	//poner datos
	bar.SetXAxis([]string{"Enero", "Febrero", "Marzo", "Abril", "Junio", "Julio"}).
		AddSeries("Category A", generateBarFile()).
		AddSeries("Category B", generateBarFile()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	bar.Render(w)
}

//Genera la grafica de lineas con archivo CSV
func lineGraphcsv(w http.ResponseWriter, _ *http.Request) {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Grafica de lineas",
			Subtitle: "Generada con un archivo CSV",
		}))
	line.SetXAxis([]string{"Enero", "Febreo", "Marzo", "Abril", "Mayo", "Junio", "Julio"}).
		AddSeries("Category A", generateLineItemscsv()).
		AddSeries("Category B", generateLineItemscsv()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}
func generateLineItemscsv() []opts.LineData {
	items := make([]opts.LineData, 0)
	sum := 0
	for i := 0; i < len(ReadCSVFile()[0]); i++ {
		sum = 0
		for j := 1; j < len(ReadCSVFile()); j++ {
			temp, _ := strconv.Atoi(ReadCSVFile()[j][i])
			sum += temp
		}
		value := strconv.Itoa(sum)
		items = append(items, opts.LineData{Value: value})
	}
	return items
}
func main() {
	op1 := os.Args[1]
	op2 := os.Args[2]

	// verificamos las opciones marcadas
	if op2 == "--generate" {
		switch op1 {
		case "--bar":
			http.HandleFunc("/", barGraph)
			http.ListenAndServe(":8080", nil)
		case "--line":
			http.HandleFunc("/", httpserver)
			http.ListenAndServe(":8080", nil)
		case "--pie":
			http.HandleFunc("/", pieGraph)
			log.Fatal(http.ListenAndServe(":8080", nil))
		default:
			log.Println("Opciones no encontradas...")
		}
	} else {

		switch op1 {
		case "--bar":
			http.HandleFunc("/", barGraphcsv)
			http.ListenAndServe(":8080", nil)
		case "--line":
			http.HandleFunc("/", lineGraphcsv)
			http.ListenAndServe(":8080", nil)

		case "--pie":
		default:
			log.Println("Opciones no encontradas...")

		}

	}
}
