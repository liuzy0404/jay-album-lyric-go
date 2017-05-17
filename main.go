package main

import (
	"fmt"
	"jay-album-lyric-go/albumId"
	"jay-album-lyric-go/request"
)

func getUserTips(slice []albumId.Album) (showUserTips string) {
	for _, album := range slice {
		showUserTips += "[" + album.ID + "] - " + album.AlbumName + "\n"
	}
	return
}

func containsString(slice []albumId.Album, input string) bool {
	for _, album := range slice {
		if input == album.ID {
			return true
		}
	}
	return false
}

func main() {
	var input string
	userTips := getUserTips(albumId.AlbumList)
	userTips += "Enter a album's id: "
	fmt.Print(userTips)
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
		return
	}
	if containsString(albumId.AlbumList, input) {
		fmt.Println("Waiting~")
		album := albumId.GetAlbumID(input)
		request.HTTPGet(request.AlbumAPI+album.AlbumID, album.AlbumName)
		fmt.Println("Done!")
		return
	}
	fmt.Println("\nYou should enter a correct album's id")
	main()
}
