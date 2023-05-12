package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(p []byte) (n int, err error) {

	for {
		b := make([]byte, 1024, 2048)
		read, err1 := rr.r.Read(b)
		if err1 == io.EOF {
			break
		}
		if err1 != io.EOF && err != nil {
			return 0, err1
		}
		for i, b2 := range b[:read] {
			p[i] = rot13byte(b2)
		}
		n += read
	}
	return n, nil
}
func rot13byte(b byte) byte {
	s := rune(b)
	if s >= 'a' && s <= 'm' || s >= 'A' && s <= 'M' {
		b += 13
	}
	if s >= 'n' && s <= 'z' || s >= 'N' && s <= 'Z' {
		b -= 13
	}

	return b
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
