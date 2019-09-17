package textoutput

import "testing"

func ExamplePrintln() {
	o := NewTextOutput(true, true)
	o.Println("hello")
	// Output:
	// hello
}

func TestColorOn(t *testing.T) {
	o := NewTextOutput(true, true)
	a := o.ColorOn(1, 36) + "x" + o.ColorOff()
	b := o.LightCyan("x")
	if a != b {
		t.Fatal(a + " != " + b)
	}
}
