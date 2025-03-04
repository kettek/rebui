package clipboard

var text string

var textGetter func() string
var textSetter func(string)

func SetTextGetter(f func() string) {
	textGetter = f
}

func SetTextSetter(f func(string)) {
	textSetter = f
}

func GetText() string {
	return textGetter()
}

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
