package iocloser

import (
	"io"
)

var errorHandler func(err error)

func Close(file io.Closer) {
	if file == nil {
		return
	}
	err := file.Close()
	if err != nil {
		if errorHandler != nil {
			errorHandler(err)
		}
	}
}
