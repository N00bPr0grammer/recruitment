package main

import (
	"fmt"
	"os"
	"encoding/csv"
	s "sort"
	q "queue"
	c "compare"
	co "concurrency"
)

func Connect() {
	res := map[string][]string{}
	json.Unmarshal(GetApi("https://data.go.id/dataset/museum-indonesia"), &res)
	for k, v := range res {
        fmt.Printf("%s=%#v\n", k, v)
    }
    csvdatafile, err := os.Create("./Kota.csv") //converting data to csv

    if err != nil {
       fmt.Println(err)
    }
    defer csvdatafile.Close()

    writer := csv.NewWriter(csvdatafile)
}

func Main() {
	Barcharts()
	Sort()
	Sort2()
	Compare()
	Connect()
}