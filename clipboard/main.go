package clipboard

var text string

var textGetter func() string
var textSetter func(string)

// SetTextGetter sets the function that is used to acquire the current clipboard contents.
func SetTextGetter(f func() string) {
	textGetter = f
}

// SetTextSetter sets the function that is used to set the current clipboard contents.
func SetTextSetter(f func(string)) {
	textSetter = f
}

// GetText returns the contents of the clipboard.
func GetText() string {
	return textGetter()
}

// SetText sets the contents of the clipboard.
func SetText(s string) {
	textSetter(s)
}

func init() {
	if textGetter == nil {
		textGetter = func() string {
			return text
		}
	}
	if textSetter == nil {
		textSetter = func(s string) {
			text = s
		}
	}
}
