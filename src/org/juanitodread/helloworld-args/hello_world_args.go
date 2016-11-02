package main

import (
  "fmt"
  "os"
)

func main() {

  fmt.Printf("The name of the command that invoked this program is: %s\n", os.Args[0])

  // Start to print since position 1, becase position 0 (zero) is reserved for
  // the name of the command that invoke the program

  // A more idiomatic for: first define two variables i and elem, i is the index
  // of the loop and elem is the element at position i of the array
  // the array starts from 1..N
  for i, elem := range os.Args[1:]{
    fmt.Printf("os.Args[%d]=%s\n", i, elem)
  }
}
