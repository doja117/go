package main

import (
	"fmt"
	"reflect"
)

type File struct {
	fd   int
	name string
}

func newFile(fd int, str string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, str}
}

func main() {

	f := newFile(10, "Saurabh")
	fmt.Printf("%v\t", reflect.TypeOf(f))
	fmt.Printf("%s\t", f.name)
}
