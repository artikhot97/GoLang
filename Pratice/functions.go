package main

import (
    "fmt"
    "strings"
)

func stringConvertUpper(str string) string {
    return strings.ToUpper(str)
}

func stringConvertLower(str string) string {
    return strings.ToLower(str)
}

func main(){
    fmt.Println("String Conversion to Upper: ", stringConvertUpper("Hello Arti..!!"))
    fmt.Println("String Conversion to Lower: ", stringConvertLower("Hello MB..!!"))
}