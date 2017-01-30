package main

import (
  "fmt"
)

func insertionSort(A []int) {
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
  a := make([]int, 6)
  a[0] = 4
  a[1] = 3
  a[2] = 1
  a[3] = 7
  a[4] = 9
  a[5] = 10

  insertionSort(a)
  
  fmt.Println(a)
}
