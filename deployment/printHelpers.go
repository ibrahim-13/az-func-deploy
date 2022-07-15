package deployment

import (
	"io"
)

const (
	__printReset string = "[-:-:-]"
)

func writeRedln(w io.Writer, msg string) {
	w.Write([]byte("[red]" + msg + __printReset + "\n"))
}

func writeHighlightln(w io.Writer, msg string) {
	w.Write([]byte("[#2dd1d6:#191bb0]" + msg + __printReset + "\n"))
}

func writeBlackYellowln(w io.Writer, msg string) {
	w.Write([]byte("[black:yellow]" + msg + __printReset + "\n"))
}

func writeGreenln(w io.Writer, msg string) {
	w.Write([]byte("[green]" + msg + __printReset + "\n"))
}
