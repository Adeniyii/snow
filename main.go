package main

import (
	"fmt"
	"os"
	"os/user"

	"snow/repl"
)

func main() {
	const welcomeAscii = `
                                   .::!!!!!!!:.
  .!!!!!:.                        .:!!!!!!!!!!!!
  ~~~~!!!!!!.                 .:!!!!!!!!!UWWW$$$
      :$$NWX!!:           .:!!!!!!XUWW$$$$$$$$$P
      $$$$$##WX!:      .<!!!!UW$$$$"  $$$$$$$$#
      $$$$$  $$$UX   :!!UW$$$$$$$$$   4$$$$$*
      ^$$$B  $$$$\     $$$$$$$$$$$$   d$$R"
        "*$bd$$$$      '*$$$$$$$$$$$o+#"
             """"          """""""
  `
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"Hi %s,\n%s\nWelcome to the snow programming language.\n", user.Username, welcomeAscii)
	repl.Start(os.Stdin, os.Stdout)
}
