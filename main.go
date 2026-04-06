package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	"github.com/xuri/excelize/v2"
)

func tranformToValue(value string) float64 {
	str := strings.ReplaceAll(strings.ReplaceAll(value, ".", ""), ",", ".")

	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}
	return res
}

func findValueByUf(uf string, file string) string {
	f, err := excelize.OpenFile(file)
	if err != nil {
		return ""
	}
	defer f.Close()
	rows, err := f.GetRows("Sheet1")

	for _, row := range rows {
		if strings.TrimSpace(row[0]) == uf {
			return row[1]
		}
	}
	return ""
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	file, err := os.Open("Lista_imoveis_geral.csv")

	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	con := os.Getenv("db")
	fmt.Println("CONNECTION STRING:", con)
	db, err := sql.Open("postgres", con)

	if err != nil {
		fmt.Println("DEU ERRO AQUI O POSSST")
		fmt.Println(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = ';'

	defer file.Close()
	counter := 0
	var totalTerrenos int

	fixUf := map[string]string{
		"AC": "Acre",
		"AL": "Alagoas",
		"AP": "Amapá",
		"AM": "Amazonas",
		"BA": "Bahia",
		"CE": "Ceará",
		"DF": "Distrito Federal",
		"ES": "Espírito Santo",
		"GO": "Goiás",
		"MA": "Maranhão",
		"MT": "Mato Grosso",
		"MS": "Mato Grosso do Sul",
		"MG": "Minas Gerais",
		"PA": "Pará",
		"PB": "Paraíba",
		"PR": "Paraná",
		"PE": "Pernambuco",
		"PI": "Piauí",
		"RJ": "Rio de Janeiro",
		"RN": "Rio Grande do Norte",
		"RS": "Rio Grande do Sul",
		"RO": "Rondônia",
		"RR": "Roraima",
		"SC": "Santa Catarina",
		"SP": "São Paulo",
		"SE": "Sergipe",
		"TO": "Tocantins",
	}

	filesToOpen := []string{"avgIdade.xlsx", "rendimento.xlsx", "sexo.xlsx", "superior.xlsx", "totalPopulacao.xlsx", "txCrescAnualPop.xlsx"}

	var usedUF []string

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

			uf := strings.TrimSpace(i[1])
			var exists bool

			for _, v := range usedUF {
				if v == uf {
					exists = true
					break
				}
			}
			if !exists {
				var allValuesByUF []string
				for _, file := range filesToOpen {
					value := findValueByUf(fixUf[strings.TrimSpace(i[1])], file)
					allValuesByUF = append(allValuesByUF, fmt.Sprintf("%.2f", tranformToValue(value)))

					// fmt.Printf("%s: %s\n", file, fmt.Sprintf("%.2f", tranformToValue(value)))
				}

				if len(allValuesByUF) == 6 {
					_, err = db.Exec(
						"INSERT INTO UF_INFO (UF, IDADE_MEDIA, RENDA_MEDIA, HOMEM_PARA_100_MULHERES, PERCENT_ENSINO_SUP, POPULACAO, TX_CRESC_ANUAL_POP) VALUES ($1, $2, $3, $4, $5, $6, $7)",
						strings.TrimSpace(i[1]),
						allValuesByUF[0],
						allValuesByUF[1],
						allValuesByUF[2],
						allValuesByUF[3],
						allValuesByUF[4],
						allValuesByUF[5],
					)
					if err != nil {
						return
					}
				}
				usedUF = append(usedUF, uf)
			}
			area := strings.Split(i[9], ",")
			_, err = db.Exec(
				"INSERT INTO IMOVEL_INFO (UF, CIDADE, BAIRRO, ENDERECO, PRECO, FINANCIAVEL, AREA) VALUES ($1, $2, $3, $4, $5, $6, $7)",
				strings.TrimSpace(i[1]),
				strings.TrimSpace(i[2]),
				strings.TrimSpace(i[3]),
				strings.TrimSpace(i[4]),
				tranformToValue(i[5]),
				strings.TrimSpace(i[8]),
				tranformToValue(strings.Split(strings.TrimSpace(area[len(area)-1]), " ")[0]),
			)
			if err != nil {
				fmt.Println("DEU ERRO AQUI O POSSST")
				fmt.Println(err)
				return
			}

		}

		counter++
	}
}
