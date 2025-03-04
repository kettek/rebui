This package provides clipboard support.

Per default the clipboard is limited to just rebui's context.

However, if the `github.com/rebui/clipboard/system` package is used as a blank import, an OS-level clipboard system will be used. This clipboard uses golang-design's clipboard which may have requirements to build depending on the OS. See [https://github.com/golang-design/clipboard?tab=readme-ov-file#platform-specific-details](here).
