package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func randArray (len int) []int {
  a := make([]int, len)
  for i := 0; i < len; i++ {
    a[i] = rand.Intn(len)
  }
  return a
}

func add (numbers []int) int64 {
  var sum int64
  for _, n := range numbers {
    sum += int64(n)
  }
  return sum
}

func addWithConcurrency (numbers []int) int64 {
  numOfCores := runtime.NumCPU()
  runtime.GOMAXPROCS(numOfCores)
  var sum int64
  max := len(numbers)
  sizeOfPart := max / numOfCores
  var wg sync.WaitGroup
  for i := 0 ; i < numOfCores; i++ {
    start := i * sizeOfPart
    end := start + sizeOfPart
    part := numbers[start:end]
    wg.Add(1)
    go func(nums []int) {
      defer wg.Done()
      var partSum int64
      for _, n := range nums {
        partSum += int64(n)
      }
      atomic.AddInt64(&sum, partSum)
    }(part)
  }

  wg.Wait()
  return sum
}

func main () {
  cpuNum := runtime.NumCPU() * 1e7
  fmt.Printf("CPU NUM %d\n", cpuNum)

  numbers := randArray(cpuNum)

  // function without concurency
  t := time.Now()
  sum := add(numbers)
  fmt.Printf("Without concurency, Sum : %d, Waktu : %s\n", sum, time.Since(t))

  // // function with concurency
  t = time.Now()
  sum = addWithConcurrency(numbers)
  fmt.Printf("With concurency, Sum : %d, Waktu : %s\n", sum, time.Since(t))
}
