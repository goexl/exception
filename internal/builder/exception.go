package builder

import (
	"github.com/goexl/exception/internal/core"
	"github.com/goexl/exception/internal/kernel"
	"github.com/goexl/exception/internal/param"
	"github.com/goexl/gox"
)

type Exception struct {
	params *param.Exception
}

func NewException() *Exception {
	return &Exception{
		params: param.NewException(),
	}
}

func (e *Exception) Code(code kernel.Code) (exception *Exception) {
	e.params.Code = code
	exception = e

	return
}

func (e *Exception) Message(message string) (exception *Exception) {
	e.params.Message = message
	exception = e

	return
}

func (e *Exception) Field(required gox.Field[any], optionals ...gox.Field[any]) (exception *Exception) {
	e.params.Data[required.Key()] = required.Value()
	for _, field := range optionals {
		e.params.Data[field.Key()] = field.Value()
	}
	exception = e

	return
}

func (e *Exception) Build() *core.Exception {
	return core.NewException(e.params)
}
