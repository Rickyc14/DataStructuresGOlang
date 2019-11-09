package main

import (
  "fmt"
)

func sort(slice []int) {
  var firstNumber, secondNumber int
  var t int = 1

  myloop:for {
    var count int = 0
    for i := 0; i < len(slice) - t; i++ {
      firstNumber = slice[i]
      secondNumber = slice[i+1]
      if firstNumber > secondNumber {
        slice[i] = secondNumber
        slice[i+1] = firstNumber
        count++
      }
    }
    if count == 0 {
      break myloop
    }
    t++
  }
}

func main() {
  xs := []int{2, 1, 9, 5, 3, 2, 3, 11, 24, 5, 67, 1, 0}
  sort(xs)
  //sort(xs)
  fmt.Println(xs)
}
