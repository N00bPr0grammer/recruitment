package concurrency

import (
	"fmt"
    	"io/ioutil"
    	"log"
    	"net/http"
   	"os"
	"encoding/csv"
)

//get api
func get_api(url string) []byte {
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    return body
}

func main() {
    res := map[string][]string{}
    json.Unmarshal(get_api("https://data.go.id/dataset/museum-indonesia"), &res)
    for k, v := range res {
        fmt.Printf("%s=%#v\n", k, v)
    }
    csvdatafile, err := os.Create("./Kota.csv") //converting data to csv

    if err != nil {
       fmt.Println(err)
    }
    defer csvdatafile.Close()

    writer := csv.NewWriter(csvdatafile)
