package controller

//声明不同任务表，以对用户做区分,目前仅有普通用户
var (
	normaltasks []int
)

//尝试构建一个词库，用于区分不同的关键词,尚在优化中
var ()

func init() {
	normaltasks = []int{1, 2, 3}
}
