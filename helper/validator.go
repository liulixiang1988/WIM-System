package helper

import (
	"regexp"
)

var DatePatten = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
