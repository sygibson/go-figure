package main

import (
	"cmd"
	"figure"
)

/* ToDo:
1. flags to represent all of the things for input
2. iterate over the list of embedded fonts
3. embed the fonts in the package
4. build up the function to write the banner
*/

/* see documentation:
	func NewFigure(phrase, fontName string, strict bool) figure
	func NewFigureWithFont(phrase string, reader io.Reader, strict bool) figure
	myFigure := figure.NewFigure("Foo Bar", "", true)
	myFigure := figure.NewFigure("Foo Bar", "alphabet", true)
	myFigure := figure.NewFigureWithFont("Foo Bar", "/home/ubuntu/go/src/github.com/common-nighthawk/go-figure/fonts/alphabet.flf", true)
	myFigure := figure.NewFigureWithFont("Foo Bar", "/home/lib/fonts/alaphabet.flf", true)
	myFigure.Blink(5000, 1000, 500)
	myFigure.Blink(5000, 1000, -1)
	myFigure.Scroll(5000, 200, "right")
  myFigure.Scroll(5000, 100, "left")
*/

func main() {
	cmd.Execute()

	myFigure := figure.NewFigure("Hello World", "", true)
	myFigure.Print()
}

/* Functions:

Blink(duration, timeOn, timeOff int)
Scroll(duration, stillness int, direction string)
Dance(duration, freeze int)


*/
