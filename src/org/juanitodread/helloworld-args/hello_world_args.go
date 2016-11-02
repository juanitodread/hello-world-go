package main

import (
  "fmt"
  "os"
)

func main() {
  var arguments string

  // Start to print since position 1, becase position 0 (zero) is reserved for
  // the name of the command that invoke the program
  for i := 1; i < len(os.Args); i++ {
    arguments += os.Args[i] + " "
    fmt.Println(i)
  }

  fmt.Printf("The name of the command that invoked this program is: %s\n", os.Args[0])
  fmt.Printf("Hello World: %s\n", arguments)
}
