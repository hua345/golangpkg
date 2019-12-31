package timeLearn

import (
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	nowTime := time.Now()
	t.Log(nowTime)
	// 时间格式化
	t.Logf("%d-%02d-%02d %02d:%02d:%02d",
		nowTime.Year(), nowTime.Month(), nowTime.Day(),
		nowTime.Hour(), nowTime.Minute(), nowTime.Second())

	// 2006-01-02 15:04:05.999999999 -0700 MST
	// 年: 06/2006表示年份
	// 月份: 数字1/01和单词Jan/January表示月份
	// 日: 2/02/_2表示几号
	// 小时: 3/03表示12小时制，15表示24小时制
	// 时区: -07/-0700/Z0700/Z07:00/-07:00/MST表示时区,-0700表示西7区，+0800表示东八区上海
	t.Log(nowTime.Format("2006-01-02 15:04:05"))
	// 时间戳/秒(10位)
	t.Log(nowTime.Unix())
	// 时间戳/纳秒(19位))
	t.Log(nowTime.UnixNano())
	// 时间戳转时间
	t.Log(time.Unix(1576554000, 0).Format("2006-01-02 15:04:05"))
}
