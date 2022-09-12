package main

import "fmt"

func main() {
  fmt.Println("hello world")
  sometihgn := Something{Name: "aaa"}
  fmt.Println(sometihgn.Name)
}

type Something struct {
  Name string
}
