package model

const (
	TYPE_JOKE = iota
	TYPE_HISTORY_TODAY
	TYPE_XIEHOUYU
	TYPE_MIYU
	TYPE_MINGRENMINGYAN
	TYPE_OTHRE
)

//News
//Title
//Content not null
//Img
//Type: 0 joke, 1 historyToday , 2 xiehouyu, 3 miyu,4 mingrenmingyan, 255 other
type News struct {
	Title   string
	Content string
	Img     string
	Type    int // 0 joke, 1 historyToday , 2 xiehouyu, 3 miyu,4 mingrenmingyan, 255 other
}

//NewNews
//Type: 0 joke, 1 historyToday , 2 xiehouyu, 3 miyu,4 mingrenmingyan, 255 other
func NewNews(title string, content string, img string, _type int) *News {
	return &News{
		Title:   title,
		Content: content,
		Img:     img,
		Type:    _type,
	}
}

//func (n *NewNews())String(){
//	fmt.Println()
//}
