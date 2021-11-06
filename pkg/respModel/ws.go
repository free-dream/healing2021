package respModel

import "time"

type SysMsg struct {
	Uid       uint      `json:"uid"`
	Type      int       `json:"type"`
	ContentId uint      `json:"contentId"`
	Song      string    `json:"song"`
	Time      time.Time `json:"time"`
	IsSend    int       `json:"isSend"`
}

type UsrMsg struct {
	FromUser uint   `json:"fromUser"`
	ToUser   uint   `json:"toUser"`
	Url      string `json:"user"` //录音url
	Song     string `json:"song"` //歌名
	Message  string `json:"message"`
	IsSend   int    `json:"isSend"`
}

type Sysmsg struct {
	Uid       uint      `json:"uid"`
	Type      int       `json:"type"`
	ContentId uint      `json:"contentId"`
	Song      string    `json:"song"`
	Time      time.Time `json:"time"`
	IsSend    int       `json:"isSend"`
}

type Usrmsg struct {
	FromUser uint   `json:"fromUser"`
	ToUser   uint   `json:"toUser"`
	Url      string `json:"user"` //录音url
	Song     string `json:"song"` //歌名
	Message  string `json:"message"`
	IsSend   int    `json:"isSend"`
}
