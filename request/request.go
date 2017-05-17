package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var (
	albumExp = regexp.MustCompile(`<textarea\s+style="display:none;">(.*)<\/textarea>`)
	lyricExp = regexp.MustCompile(`\[.*\]`)
	ua       = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.36 Safari/537.36"
	lyricAPI = "http://music.163.com/api/song/media?id="
	// AlbumAPI export albumApi
	AlbumAPI = "http://music.163.com/album?id="
)

type song struct {
	Name string
	ID   int
}

type lyric struct {
	Lyric string
}

// HTTPGet export HTTPGet func
func HTTPGet(url, albumName string) (content string) {
	var songList []song
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error.", err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("http read error.", err)
		return
	}
	// get matched array
	submatch := albumExp.FindSubmatch([]byte(data))
	content = string(submatch[1])
	json.Unmarshal([]byte(content), &songList)
	for _, song := range songList {
		//fmt.Println(song.Name, "'s id is ", song.ID)
		getLyric(lyricAPI+strconv.Itoa(song.ID), albumName, song.Name)
	}
	return
}

func getLyric(url, albumName, songName string) {
	var lyricContent lyric
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", ua)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error.", err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("http read error.", err)
		return
	}

	content := string(data)
	json.Unmarshal([]byte(content), &lyricContent)
	lyricContent.Lyric = lyricExp.ReplaceAllString(lyricContent.Lyric, "")
	fileName := albumName + "/" + songName + ".txt"
	// check whether album direction exist or not
	if _, err := os.Stat("./" + albumName); os.IsNotExist(err) {
		os.Mkdir(albumName, 0777)
	}
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("create file failed.", err)
		return
	}
	file.Write([]byte(lyricContent.Lyric))
	fmt.Println(songName + " saved!")
}
