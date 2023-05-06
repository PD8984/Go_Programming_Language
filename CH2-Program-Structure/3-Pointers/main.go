package main

import "fmt"

func number(i int) {
    i = 0
}

func numptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial Value:", i)

    number(i)
    fmt.Println("number:", i)

    numptr(&i)
    fmt.Println("numptr:", i)

    fmt.Println("pointer:", &i)
}
