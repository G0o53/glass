package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("aaaaaaaaaaa")
    return
  }
  filename := os.Args[1]

  file, err := os.Open(filename)
  if err != nil {
    fmt.Println("file failed to open, does it exist?")
  }

  defer file.Close()
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()

  }
}

