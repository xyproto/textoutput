package main

import (
	"github.com/xyproto/textoutput"
)

func main() {
	// Prepare to output colored text, but respect the NO_COLOR environment variable
	o := textoutput.New()

	o.Print("<blue>hi</blue> ")
	o.Println("<yellow>there</yellow>")
	o.Printf("<green>%s: <red>%d<off>\n", "number", 42)
}
