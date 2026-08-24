package collections

import (
	"bufio"
	"io"
)

// newLineScanner bounds line length so a bad upload can not eat memory.
func newLineScanner(r io.Reader) *bufio.Scanner {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 64*1024), 512*1024)
	return sc
}
