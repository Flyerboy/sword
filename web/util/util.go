package util

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"regexp"
	"strings"
)

func MarkDown(src string) (dst string)  {
	src = strings.Replace(src, "\r\n", "\n", -1)
	unsafe := blackfriday.Run([]byte(src))
	p := bluemonday.UGCPolicy()
	p.AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("pre")
	html := p.SanitizeBytes(unsafe)
	return string(html)
}