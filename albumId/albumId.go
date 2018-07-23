package albumId

// Album export album struct
type Album struct {
	ID        string
	AlbumID   string
	AlbumName string
}

// AlbumList export album list
var AlbumList = []Album{
	{"1", "18918", "jay"},
	{"2", "18915", "范特西"},
	{"3", "18907", "八度空间"},
	{"4", "18905", "叶惠美"},
	{"5", "18903", "七里香"},
	{"6", "18896", "十一月的萧邦"},
	{"7", "18893", "依然范特西"},
	{"8", "18886", "我很忙"},
	{"9", "18877", "魔杰座"},
	{"10", "18875", "跨时代"},
	{"11", "18869", "惊叹号"},
	{"12", "2263029", "十二新作"},
	{"13", "3084335", "哎呦，不错哦"},
	{"14", "34720827", "周杰伦的床边故事"},
	{"15", "37251353", "等你下课"},
	{"16", "38721188", "不爱我就拉倒"},
}

// GetAlbumID for get album's id from user's input
func GetAlbumID(idx string) (aAlbum Album) {
	for _, album := range AlbumList {
		if idx == album.ID {
			aAlbum.AlbumID = album.AlbumID
			aAlbum.AlbumName = album.AlbumName
		}
	}
	return
}
