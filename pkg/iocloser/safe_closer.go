package iocloser

import (
	"io"
)

var errorHandler func(err error)

func Close(file io.Closer) {
	err := file.Close()
	if err != nil {
		if errorHandler != nil {
			errorHandler(err)
		}
	}
}
