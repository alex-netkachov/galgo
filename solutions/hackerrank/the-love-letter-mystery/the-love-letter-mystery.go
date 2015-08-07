package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

// https://www.hackerrank.com/challenges/the-love-letter-mystery

func main() {
  s := bufio.NewScanner(bufio.NewReader(os.Stdin))

  s.Scan()
  n, _ := strconv.Atoi(s.Text())

  for i := 0; i < n; i++ {
    s.Scan()
    fmt.Printf("%v\n", countStepsForPalindrome(s.Text()))
  }
}

func countStepsForPalindrome(text string) int {
  sum := 0
  length := len(text)
  half := length / 2
  for i := 0; i < half; i++ {
    v := int(text[i]) - int(text[length-i-1])
    if v < 0 {
      v = -v
    }
    sum += v
  }
  return sum
}
