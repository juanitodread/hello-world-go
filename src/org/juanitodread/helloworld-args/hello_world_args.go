package main

import (
  "fmt"
  "os"
)

func main() {
  var arguments string

  for i := 1; i < len(os.Args); i++ {
    arguments += os.Args[i] + " "
  }

  fmt.Printf("Hello World: %s\n", arguments)
}
