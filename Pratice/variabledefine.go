package main

import "fmt"


func main(){
    var a int = 10
    b := 5

    const (
        c = "Hello"
    )

    var s string = "Ok Google"
    fmt.Println(s)
    d := []rune(s) // Rune is an alias for int32
    fmt.Println(d)
    d[0] = 'B'
    d[1] = 'y'
    fmt.Println(d)
    s2 := string(d)

    //c := "Test"

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(s2)
}