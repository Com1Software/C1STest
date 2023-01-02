package main

import (
	"github.com/Com1Software/C1S"
	"fmt"
	"os"
	"runtime"
	term "github.com/nsf/termbox-go"
)

func main() {
	fmt.Printf("C1Script\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
//	args := os.Args
	err := term.Init()
	cbuf := ""
	if err != nil {
		panic(err)
	}
	defer term.Close()
	switch {
	//-------------------------------------------------------------
	case len(os.Args) == 2:
	   fmt.Printf("cmd %s\n",os.Args[1])
	   c1s.Cmd(os.Args[1])
	fmt.Println("Not")

		//}
		//-------------------------------------------------------------
	default:
		cmd := ""
		fmt.Printf("C1Script\n")
		fmt.Printf("Operating System : %s\n", runtime.GOOS)
		fmt.Println("Command Line Mode")
		fmt.Printf(".")
		for {
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {
				case term.KeyArrowUp:
					fmt.Println("Arrow Up pressed")
				case term.KeyArrowDown:
					fmt.Println("Arrow Down pressed")
				case term.KeyEnter:
					fmt.Printf("\n.")
					cmd = cbuf
					cbuf = ""
				default:
					s := rune(ev.Ch)
					cbuf = cbuf + string(s)
					fmt.Printf("%s", string(s))
				}
			}
			switch {
			case cmd == "exit":
				fmt.Println("\nExit C1Script")
				os.Exit(2)
			case cmd == "?":
				fmt.Println("\nHelp")

				fmt.Printf("\n.")
			}
		}
	}

}
