package param

import (
	"github.com/goexl/exception/internal/internal"
)

type Exception struct {
	Code    int             `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	Data    internal.Fields `json:"data,omitempty"`
}

func NewException() *Exception {
	return &Exception{
		Data: make(map[string]any),
	}
}
