package respModel

import "time"

type SysMsg struct {
	Uid       uint      `json:"uid"`
	Type      int       `json:"type"`
	ContentId uint      `json:"contentId"`
	Song      string    `json:"song"`
	Time      time.Time `json:"time"`
	IsSend    int       `json:"isSend"`
    FromUser  string    `json:"fromUser"`
}

type UsrMsg struct {
	FromUser  uint      `json:"fromUser"`
	ToUser    uint      `json:"toUser"`
    FromUserName string `json:"fromUserName"`
    ToUserName string   `json:"toUserName"`
	Url       string    `json:"user"` //录音url
	Song      string    `json:"song"` //歌名
	SongId    uint      `json:"songId"`
	Message   string    `json:"message"`
	IsSend    int       `json:"isSend"`
	CreatedAt time.Time `json:"time"`
}

type Sysmsg struct {
	Uid       uint      `json:"uid"`
	Type      int       `json:"type"`
	ContentId uint      `json:"contentId"`
	Song      string    `json:"song"`
	Time      time.Time `json:"time"`
	IsSend    int       `json:"isSend"`
    FromUser  string    `json:"fromUser"`
}

type Usrmsg struct {
	FromUser  uint      `json:"fromUser"`
	ToUser    uint      `json:"toUser"`
    FromUserName string `json:"fromUserName"`
    ToUserName string   `json:"toUserName"`
	Url       string    `json:"user"` //录音url
	Song      string    `json:"song"` //歌名
	SongId    uint      `json:"songId"`
	Message   string    `json:"message"`
	IsSend    int       `json:"isSend"`
	CreatedAt time.Time `json:"time"`
}
