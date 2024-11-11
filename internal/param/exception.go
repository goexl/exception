package param

import (
	"github.com/goexl/exception/internal/internal"
	"github.com/goexl/exception/internal/kernel"
)

type Exception struct {
	Code    kernel.Code     `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	Data    internal.Fields `json:"data,omitempty"`
}

func NewException() *Exception {
	return &Exception{
		Data: make(map[string]any),
	}
}
