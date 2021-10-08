package main

import (
	"fmt"
	"io"
	"strings"
)

/* io.Reader interface
type Reader interface {
	Read(p []byte) (n int, err error)
}
*/

func _012Readers() {
	newReader := strings.NewReader("Hello, world!")

	buf := make([]byte, 8)

	for {
		bytesRead, err := newReader.Read(buf)
		fmt.Printf("bytesRead = %v err = %v buf = %v\n", bytesRead, err, buf)
		fmt.Printf("buf[:bytesRead] = %q\n", buf[:bytesRead])
		if err == io.EOF {
			break
		}
	}
}
