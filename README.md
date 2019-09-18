# textoutput

Package for controlling text output, with or without colors, on Linux, using VT100 terminal codes.

## Example use

```go
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

	// Output "a" in light blue and "b" in light green
	fmt.Println(o.Words("a b", "blue", "green"))

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
```

## General info

* Version: 1.6.0
* License: MIT
* Author &lt;xyproto@archlinux.org&gt;
