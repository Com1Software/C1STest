package main

import (
	"fmt"
	"os"
	"runtime"

	c1s "github.com/Com1Software/C1S"
	term "github.com/nsf/termbox-go"
)

func main() {
	cmdctl := c1s.CommandControl{}
	fmt.Printf("C1Script\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	//	args := os.Args
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	switch {
	//-------------------------------------------------------------
	case len(os.Args) == 2:
		fmt.Printf("cmd %s\n", os.Args[1])
		c1s.Run(os.Args[1])
		fmt.Println("Not")

		//}
		//-------------------------------------------------------------
	default:
		cmdctl.Command = "test"
		cmdctl = c1s.Cmd(cmdctl)
		fmt.Printf("Command %s\n", cmdctl.Command)

	}
}
