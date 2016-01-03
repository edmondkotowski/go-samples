package main

import (
    "fmt"
)

type Foo interface {
    Run() string
}

type Bar struct {
}

func (b Bar) Run() string {
    return "Bar"
}

type Baz struct {
}

func (b *Baz) Run() string {
    return "Dynamic Baz"
}

func main() {
    test_objects := []Foo{Bar{}, new(Baz)}
    for _, object := range test_objects {
        fmt.Println(object.Run())
    }
}