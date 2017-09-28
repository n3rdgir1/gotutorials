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

// global variables go here, get shared with everything in the package if it has a lowercase level
// if you want variable to be available to EVERYONE!!!!, start it with a capital letter
//above rules apply to method names, etc, as well

func main() {
  // local variables go here
  lang := "English" // `:=` initialization and assignment at the same time, can only be used inside functions

  if len(os.Args) >= 2 {
    lang = os.Args[1]
  }
  fmt.Println(greetlib.GreetIn(lang))
}

// _ is a thing in go, necessary since you can't have unused variables (compiler enforced)

/*
 if you don't consciously handle the error, nothing will. no try, catch.
 If there are multiple return values, the last one might be an error, might also be `nil`. WRITE YOUR DOCS!! (godocs)
 */
