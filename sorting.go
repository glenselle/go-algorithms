package main

import (
  "fmt"
)

func selectionSort(A []int) {
  for j := 0; j < len(A); j++ {
    key := A[j]
    m := j

    for i := j; i < len(A); i++ {
      if A[i] < A[m] {
        m = i
      }
    }

    A[j] = A[m]
    A[m] = key
  }
}

func main() {
  a := []int{4, 3, 1, 9, 7, 10} // [4 3 1 9 7 10]

  selectionSort(a)
  
  fmt.Println(a) // [1 3 4 7 9 10]
}
