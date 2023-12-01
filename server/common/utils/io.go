package utils

import "io"

func CloseQuiet(c io.Closer) {
	_ = c.Close()
}
