package concurrency

import (
    "io/ioutil"
    "net/http"
)

//get api
func GetApi(url string) []byte {
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
