package utils

import "regexp"

var cropRgx *regexp.Regexp

// StrCrop crops a string like you want in blog previews
func StrCrop(str string, length int) string {
	if len(str) <= length {
		return str
	}

	if cropRgx == nil {
		cropRgx = regexp.MustCompile(`[\s,.-]`)
	}

	var (
		lastStop int
		runeStr  = []rune(str)
	)
	for i, r := range runeStr {
		if cropRgx.MatchString(string(r)) {
			lastStop = i
		}
		if i >= length {
			return string(runeStr[:lastStop])
		}
	}

	return str
}
