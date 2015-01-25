package main

import "os"
import "bufio"

func main() {
  stdin := bufio.NewReader(os.Stdin)
  stdout := bufio.NewWriter(os.Stdout)

  defer stdout.Flush()

  for {
    line, _, err := stdin.ReadLine()

    if err != nil {
      break
    }

    stdout.Write(append(line, '\n'))
  }
}
