package main

import (
	"fmt"
	"github.com/xyproto/textoutput"
)

func main() {
	// Enable colors, enable output
	o := textoutput.NewTextOutput(true, true)

	s := o.Tags("<red>hi</red> and <lightblue>hello</lightblue>")

	fmt.Println(s)

	for _, cc := range o.Extract(s) {
		fmt.Printf("letter: %d\tcolor attributes: %v\n", cc.R, cc.A.Ints())
	}
}
