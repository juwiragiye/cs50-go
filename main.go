package main

import (
	"fmt"

	"github.com/juwiragiye/cs50-go/sql"
)

func main() {
	data := sql.CountMovies("favorites.csv")
	for title, count := range data {
		fmt.Println(title, count)
	}


	
}	