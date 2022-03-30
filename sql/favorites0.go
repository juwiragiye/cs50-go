package sql

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

type Movie struct {
	title string
	count int
}

func ReadCSV(filename string) ([]string, error) {
	file, err  := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	var favorites []string
	

	if err != nil {
		log.Fatal(err)
	}

	for {
        rec, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)

        }
        // do something with read line
        favorites = append(favorites, strings.ToUpper(string((rec[1]))))
    
	}

	return favorites, nil

}

func FormatCSV(filename string) []string {
	data, err := ReadCSV(filename)
	var formatedCSV []string

	if err != nil {
		log.Fatal(err)
	}

	for _, title := range data {
		if title == "TITLE" {
			continue
		}
		formatedCSV = append(formatedCSV, strings.Trim(title, "\n"))
	}

	return formatedCSV
}

 
func CountMovies(filename string) map[string]int {
	moviesCount := make(map[string]int)
	formattedData := FormatCSV(filename)

	for _, title := range formattedData {
		_, ok := moviesCount[title]
		if ok  {
			moviesCount[title]++
		} else {
			moviesCount[title] = 1
		}
	}

	return moviesCount
}
