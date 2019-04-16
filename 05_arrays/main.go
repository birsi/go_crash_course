package main

import "fmt"

func main() {
  // var fruitArr[2] string
  //
  // fruitArr[0] = "Apple"
  // fruitArr[1] = "Banana"

  fruitArr := [2]string{"Apple", "Banana"}

  fmt.Println(fruitArr)
  fmt.Println(len(fruitArr))
  fmt.Println(fruitArr[0:1])
}
