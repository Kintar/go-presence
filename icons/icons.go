package icons

import (
	_ "embed"
)

//go:embed working.png
var Working []byte

//go:embed waiting.png
var Waiting []byte

//go:embed pause.png
var Pause []byte

//go:embed quit.png
var Quit []byte
