package albumId

type album struct {
	id        string
	albumID   string
	albumName string
}

type albums []album

var albumList = albums{
	{"1", "18918", "jay"},
	{"2", "18915", "范特西"},
	{"3", "18907", "八度空间"},
	{"4", "18905", "叶惠美"},
	{"5", "18918", "七里香"},
	{"6", "18896", "十一月的萧邦"},
	{"7", "18893", "依然范特西"},
	{"8", "18886", "我很忙"},
	{"9", "18877", "魔杰座"},
	{"10", "18875", "跨时代"},
	{"11", "18869", "惊叹号"},
	{"12", "2263029", "十二新作"},
	{"13", "3084335", "哎呦，不错哦"},
	{"14", "34720827", "周杰伦的床边故事"},
}

func GetAlbumID(idx string) interface{} {
	for _, album := range albumList {
		if idx == album.albumID {
			mAlbum := map[string]string{"albumID": album.albumID, "albumName": album.albumName}
			return mAlbum
		}
	}
	return "Choose a correct album."
}
