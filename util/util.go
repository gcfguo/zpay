package util

import (
	"encoding/json"
	"unsafe"
)

func JSON(o any) string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}
