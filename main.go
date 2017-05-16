package main

import (
	"fmt"
	"jay-album-lyric-go/request"
)

func main() {
	request.HttpGet("http://music.163.com/album?id=18918")
	fmt.Println("ok")
	/*content, code := request.HttpGet("http://music.163.com/album?id=18918")
	fmt.Println(content)
	fmt.Println("code is ", code)*/
}
