package main

import "fmt"

func main() {
    fmt.Println("TeamCity Go Project")
    fmt.Println("Build successful!")
    
    a, b := 5, 3
    sum := Add(a, b)
    fmt.Printf("%d + %d = %d\n", a, b, sum)
}


func Add(x, y int) int {
    return x + y
}
