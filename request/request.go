package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jay-album-lyric-go/utils"
	"os"
	"regexp"
	"time"
)

var (
	albumExp       = regexp.MustCompile(`<textarea id="song-list-pre-data" style="display:none;">(.*)<\/textarea>`)
	lyricExp       = regexp.MustCompile(`\[.*\]`)
	ua             = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.36 Safari/537.36"
	mobileUa       = "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"
	lyricAPI       = "http://music.163.com/api/song/media?id="
	AlbumAPI       = "http://music.163.com/album?id="
	MobileAlbumAPI = "https://music.163.com/m/album?id="
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
	lyricChan := make(chan string)
	req, client := utils.HTTPClient(url, ua)
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
	decompressData, err := utils.ParseGzip(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// get matched array
	submatch := albumExp.FindSubmatch([]byte(decompressData))
	content = string(submatch[1])
	json.Unmarshal([]byte(content), &songList)
	for _, song := range songList {
		tStart := time.Now().UnixNano() / 1e6
		//fmt.Println(song.Name, "'s id is ", song.ID)
		go getLyric(lyricAPI+fmt.Sprintf("%d", song.ID), lyricChan)

		// check whether album direction exist or not
		if _, err := os.Stat("./" + albumName); os.IsNotExist(err) {
			os.Mkdir(albumName, 0777)
		}

		fileName := albumName + "/" + song.Name + ".txt"
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("create file failed.", err)
			return
		}
		file.Write([]byte(<-lyricChan))
		tEnd := time.Now().UnixNano() / 1e6
		t := fmt.Sprintf("%d", tEnd-tStart)
		fmt.Println(song.Name + " saved in " + t + "ms")
	}
	close(lyricChan)
	return
}

func getLyric(url string, lyricChan chan string) {
	var lyricContent lyric
	req, client := utils.HTTPClient(url, ua)
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
	unzipedData, err := utils.ParseGzip(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(unzipedData)
	json.Unmarshal([]byte(content), &lyricContent)
	lyricContent.Lyric = lyricExp.ReplaceAllString(lyricContent.Lyric, "")
	lyricChan <- lyricContent.Lyric
}
