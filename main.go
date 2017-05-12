package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	albumFile := "album.json"
	buf, err := ioutil.ReadFile(albumFile)
	if err != nil {
		fmt.Printf("File err: %s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
}
