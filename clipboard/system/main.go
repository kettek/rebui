package system

import (
	"fmt"

	clipboard "github.com/kettek/rebui/clipboard"
	xclipboard "golang.design/x/clipboard"
)

func init() {
	if err := xclipboard.Init(); err != nil {
		fmt.Println(err)
	} else {
		clipboard.SetTextGetter(func() string {
			return string(xclipboard.Read(xclipboard.FmtText))
		})
		clipboard.SetTextSetter(func(s string) {
			xclipboard.Write(xclipboard.FmtText, []byte(s))
		})
	}
}
