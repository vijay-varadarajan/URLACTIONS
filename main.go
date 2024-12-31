package main

import (
	"fmt"
)

func main() {
	cmdFlags := newCmdFlag()
	cmdFlags.execute()
	fmt.Println("URL-Action successful.")
}