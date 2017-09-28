package main
  /* every source code must declare a package
   * main is a special package.
   * Must live with "func main"
   */

import "fmt" //fmt is formatted input-output

func main() { //main, standard fair for a "main" function.
  fmt.Println("Hello, World!")
}

/*
pwd: ~/go/src/foo, import_path: foo, default_name: foo
pwd: ~/go/src/foo/bar, import_path: foo/bar, default_name: bar
pwd: ~/go/src/foo/bar/bazz, import_path: foo/bar/bazz, default_name: bazz

override default name:
  import ba "foo/bar/bazz" => now it's ba
*/
