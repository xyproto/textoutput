package main

import (
	"fmt"
	"github.com/xyproto/textoutput"
)

func main() {
	// Enable colors, enable output
	o := textoutput.NewTextOutput(true, true)

	// Output "a" in light blue and "b" in light green
	fmt.Println(o.LightTags("<blue>", "a", "<off> <green>", "b", "<off>"))

	// Output "a" in light blue and "b c" in light green
	fmt.Println(o.Words("a b c", "blue", "green"))

	// Output "a" in light blue and "b" in light green
	fmt.Println(o.LightBlue("a") + " " + o.LightGreen("b"))

	// Output "c" in dark blue and "d" in light yellow
	fmt.Println(o.DarkTags("<blue>c</blue> <lightyellow>d<off>"))

	// Output "a" in light blue
	fmt.Println(o.LightTags("<blue>", "a", "</blue>"))

	// Output "a" in light blue
	o.Tags("<blue>a</blue>")

	// Output "a" in light blue
	o.Tags("<blue>a<off>")

	// Exit with a dark red error message
	o.ErrExit("error: too convenient")
}
