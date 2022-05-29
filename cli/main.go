package main

import (
	"fmt"
	"log"
	"os"
	"simplifica/src/domain"
	"simplifica/src/support"
	"strings"

	"code.sajari.com/docconv"
)

func main() {
	dir, err := os.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dir {
		if strings.Contains(file.Name(), ".pdf") {
			res, err := docconv.ConvertPath(file.Name())

			if err != nil {
				log.Fatalf("Error converting %s: %s", file.Name(), err)
			}

			lines := strings.Split(res.Body, "\n")

			stckn := support.FilterLinesByText(lines, "11 - ")
			stckd := support.FilterLinesByText(lines, "R$")

			info := domain.CollectStocks(stckn, stckd)

			stocks := domain.BuildStocks(info)

			fmt.Println(stocks)
		}
	}
}
