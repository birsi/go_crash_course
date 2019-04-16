package main

import "fmt"

func greeting(name string) string {
  return "Hello " + name
}

func getSum(x, y int) int {
  return x + y
}

func main() {
  fmt.Println(greeting("Michael"))
  fmt.Println(getSum(1,2))
}
