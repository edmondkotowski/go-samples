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

// Embedded Types
type UberFoo struct {
    Bar  // Compostion
    Name string
}

func (uf *UberFoo) Sprint() error {
    fmt.Printf("Calling inner run %s with name %s\n", uf.Run(), uf.Name)

    return nil
}

func main() {
    test_objects := []Foo{Bar{}, new(Baz)}
    for _, object := range test_objects {
        fmt.Println(object.Run())
    }

    uberfoo := UberFoo{Name: "UberFoo"}
    uberfoo.Sprint()
}