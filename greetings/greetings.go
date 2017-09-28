package main

import (
  "fmt"
  "os"
  "gotutorials/greetlib"
)

/*
  When you are importing things, you can use
  `go get <url without protocol pointing to the package location>`
  to get a package off github
  (or some other version-controlled source outside my box)
  to go pull it in like a true lazy bastard :thumbsup:
*/

func main() {
  lang := "English"
  if len(os.Args) >= 2 {
    lang = os.Args[1]
  }
  fmt.Println(greetlib.GreetIn(lang))
}
