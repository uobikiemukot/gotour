package main

import (
	"io"
	"os"
	"strings"
	//"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	var count int
	var err error
	c := make([]byte, 32)

	for {
		n, err := rot.r.Read(c)
		if err != nil {
			return count, err
		}
		count += n

		for i, v := range c[:n] {
			switch {
			case ('A' <= v && v <= 'M') || ('a' <= v && v <= 'm'):
				b[i] = v + 13
			case ('N' <= v && v <= 'Z') || ('n' <= v && v <= 'z'):
				b[i] = v - 13
			default:
				b[i] = v
			}
		}
	}
	return count, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
