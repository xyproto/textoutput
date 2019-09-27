package main

import (
	"fmt"
	"github.com/xyproto/textoutput"
)

func main() {
	// Enable colors, enable output
	o := textoutput.NewTextOutput(true, true)

	s := o.Tags("<red>hi</red> and <blue>hello</blue>")

	fmt.Println(o.Extract(s))
}
