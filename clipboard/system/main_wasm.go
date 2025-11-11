//go:build wasm

package system

import (
	clipboard "github.com/kettek/rebui/clipboard"
	"syscall/js"
)

func init() {
	if js.Global().Get("navigator").Get("clipboard").Get("readText").IsUndefined() {
		js.Global().Get("alert").Invoke("Clipboard API is not supported in this browser.")
		return
	}

	clipboard.SetTextGetter(func() string {
		wait := make(chan string)
		js.Global().Get("navigator").Get("clipboard").Call("readText").Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
			wait <- args[0].String()
			return nil
		}))
		return <-wait
	})
	clipboard.SetTextSetter(func(s string) {
		wait := make(chan any)
		js.Global().Get("navigator").Get("clipboard").Call("writeText", s).Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
			wait <- nil
			return nil
		}))
		<-wait
	})
}
