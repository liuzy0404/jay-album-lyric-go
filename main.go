package main

import (
	"fmt"
	"jay-album-lyric-go/request"
)

func main() {
	request.HttpGet("http://music.163.com/album?id=18918", "jay")
	fmt.Println("ok")
}
