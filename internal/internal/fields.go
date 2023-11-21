package internal

import (
	"fmt"

	"github.com/goexl/exception/internal/internal/constant"
	"github.com/goexl/gox"
)

type Fields map[string]any

func (f Fields) String() string {
	builder := gox.StringBuilder()
	builder.Append(constant.JsonStart)
	for key, value := range f {
		builder.Append(fmt.Sprintf(`%s = %v`, key, value))
	}
	builder.Append(constant.JsonEnd)

	return builder.String()
}
