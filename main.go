package main

import (
	"fmt"
	"github.com/kpfaulkner/grafanadatadog/pkg"
)




func main() {
	fmt.Printf("so it begins....\n")
  s := pkg.NewServer()

  s.Run()
}
