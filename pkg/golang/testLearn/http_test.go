package testLearn

// Sub Test
// --------

// Sub test helps us streamline our test functions, filters out command-line level big tests into
// smaller sub tests.

import (
	"net/http"
	"testing"
)

// TestSub validates the http Get function can download content and
// handles different status conditions properly.
func TestHttp(t *testing.T) {
	tests := []struct {
		Name       string
		Url        string
		StatusCode []int
	}{
		{"statusOk", "https://www.baidu.com/", []int{http.StatusOK}},
		{"statusOk", "https://www.cnblogs.com/test.html", []int{http.StatusNotFound, http.StatusBadRequest}},
	}
	for index, item := range tests {
		t.Logf("\t单元测试: %d\t请求 %q 预计结果 %d", index, item.Url, item.StatusCode)
		{
			resp, err := http.Get(item.Url)
			if err != nil {
				t.Fatalf("请求失败 : %v", err)
			}
			t.Logf("请求成功！")

			defer resp.Body.Close()

			var requestOk bool = false

			for _, status := range item.StatusCode {
				if resp.StatusCode == status {
					requestOk = true
					break
				}
			}
			if requestOk {
				t.Logf("实际结果状态: %d ", resp.StatusCode)
			} else {
				t.Errorf("预计结果状态: %d 实际结果状态: %v", item.StatusCode, resp.StatusCode)
			}
		}
	}
}
