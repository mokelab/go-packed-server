package data

import (
	"encoding/base64"
)

type Entry struct {
	src         string
	raw         []byte
	ContentType string
}

type Cache struct {
	Entry map[string]Entry
}

func (e *Entry) Get() []byte {
	if e.raw != nil {
		return e.raw
	}
	e.raw, _ = base64.StdEncoding.DecodeString(e.src)
	return e.raw
}
