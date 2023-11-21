package exception

import (
	"github.com/goexl/exception/internal/builder"
)

func New() *builder.Exception {
	return builder.NewException()
}
