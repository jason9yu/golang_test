package main

import (
    "fmt"
)

func main() {
      var a string = "hello \n yu"
      var b string = "world"
      var c string = "Ȼ"
      
      fmt.Println(a)
      fmt.Println("----------------------------------")
      fmt.Println(len(a))
      fmt.Println(len(c))
      fmt.Println(a[0],a[1])
      fmt.Println(a + b)
}
