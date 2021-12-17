package static

import (
	_ "embed"
)

//go:embed banner.txt
var Banner string

//go:embed cfg-tamper.ini
var Config string




