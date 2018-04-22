package main

import "fmt"
import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	//b = append(b, 'A')
	b[0] = 'A'
	return 1, nil
}

func main() {
	var r MyReader
	b := make([]byte, 8)

	fmt.Println(r.Read(b))
	fmt.Printf("%q\n", b)

	reader.Validate(MyReader{})
}
