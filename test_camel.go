package main

import (
    "github.com/iancoleman/strcase"
    "fmt"
)

func main() {
    s := "andaman and nicobar"
    res := strcase.ToCamel(s)
    
    fmt.Println(res)
}
