package main

import (
    "net/http"
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
    resp, err := http.Get("https://www.reddit.com/.json")

    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }

    defer resp.Body.Close()
    contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }

    fmt.Printf("%s\n", string(contents))
}