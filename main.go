package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("Lista_imoveis_geral.csv")

	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = ';'

	defer file.Close()
	counter := 0
	var totalTerrenos int

	for {
		i, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		if strings.Contains(i[9], "Terreno") {
			totalTerrenos++
		}

		counter++
	}
	fmt.Println("Total de terrenos encontrados: ", totalTerrenos)

}
