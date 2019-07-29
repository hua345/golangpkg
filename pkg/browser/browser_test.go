package browser

import "testing"
import "github.com/pkg/browser"

func TestBrowser(t *testing.T) {
	const url = "http://192.168.137.128:8500/ui"
	err := browser.OpenURL(url)
	if err != nil {
		panic(err)
	}
}
