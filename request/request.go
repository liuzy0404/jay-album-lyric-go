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
	lyricPAI = "http://music.163.com/api/song/media?id="
)

type song struct {
	Name string
	Id   int
}

type lyric struct {
	Lyric string
}

func HttpGet(url, albumName string) (content string, statuscode int) {
	var songList []song
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error.")
		statuscode = -100
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("http read error.")
		statuscode = -200
		return
	}
	statuscode = res.StatusCode
	// get matched array
	submatch := albumExp.FindSubmatch([]byte(data))
	content = string(submatch[1])
	json.Unmarshal([]byte(content), &songList)
	for _, song := range songList {
		//fmt.Println(song.Name, "'s id is ", song.Id)
		fmt.Println("[Get] " + lyricPAI + strconv.Itoa(song.Id))
		getLyric(lyricPAI+strconv.Itoa(song.Id), albumName, song.Name)
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
		fmt.Println("http get error.")
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("http read error.")
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
		fmt.Println("create file failed", err)
		return
	}
	file.Write([]byte(lyricContent.Lyric))
}
