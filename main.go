package main

import (
  "os"
  "bufio"
  "github.com/koron/beni"
)

func main() {
  stdin := bufio.NewReader(os.Stdin)
  stdout := bufio.NewWriter(os.Stdout)
  defer stdout.Flush()

  beni.Highlight(stdin, stdout, "go", "base16", "Terminal256")
}
