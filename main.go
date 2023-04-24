package main

import (
	"fmt"
	"os"

	"github.com/dvkond/nimitta-choodamani/nimittachoodamanicore"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(nimittachoodamanicore.About())
		fmt.Println(nimittachoodamanicore.Help())
	} else {
		var dice1, dice2, dice3 = nimittachoodamanicore.RollTheDice()
		fmt.Printf("dice [%s][%s][%s]\n", dice1, dice2, dice3)
		fmt.Printf("answer = %s\n", nimittachoodamanicore.Answer(dice1, dice2, dice3))
	}
}
