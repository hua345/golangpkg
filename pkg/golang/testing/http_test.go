package testing

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
		StatusCode int
	}{
		{"statusOk", "https://www.baidu.com/", http.StatusOK},
		{"statusOk", "https://www.cnblogs.com/test.html", http.StatusBadRequest},
	}
	for index, item := range tests {
		t.Logf("\tTest: %d\tWhen checking %q for status code %d", index, item.Url, item.StatusCode)
		{
			resp, err := http.Get(item.Url)
			if err != nil {
				t.Fatalf("Should be able to make the Get call : %v", err)
			}
			t.Logf("Should be able to make the Get call.")

			defer resp.Body.Close()

			if resp.StatusCode == item.StatusCode {
				t.Logf("Should receive a %d status code.", item.StatusCode)
			} else {
				t.Errorf("Should receive a %d status code : %v", item.StatusCode, resp.StatusCode)
			}
		}
	}
}
