package main

import (
  "fmt"
  "math"
  "github.com/birsi/03_packages/strutil"
)

func main()  {
  fmt.Println(math.Floor(2.45))
  fmt.Println(math.Ceil(2.7))
  fmt.Println(math.Sqrt(16))

  fmt.Println(strutil.Reverse("birsi"))
}
