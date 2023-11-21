package core

import (
	"encoding/json"

	"github.com/goexl/exception/internal/internal/constant"
	"github.com/goexl/exception/internal/param"
	"github.com/goexl/gox"
)

type Exception struct {
	params *param.Exception
}

func NewException(params *param.Exception) *Exception {
	return &Exception{
		params: params,
	}
}

func (e *Exception) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.params)
}

func (e *Exception) Error() (error string) {
	if bytes, err := e.MarshalJSON(); nil != err {
		error = e.error()
	} else {
		error = string(bytes)
	}

	return
}

func (e *Exception) String() (str string) {
	if bytes, err := e.MarshalJSON(); nil != err {
		str = e.error()
	} else {
		str = string(bytes)
	}

	return
}

func (e *Exception) error() string {
	builder := gox.StringBuilder()
	builder.Append(constant.JsonStart)
	if 0 != e.params.Code {
		builder.Append(constant.Code).Append(constant.Equal).Append(e.params.Code)
	}
	if "" != e.params.Message {
		builder.Append(constant.Message).Append(constant.Equal).Append(e.params.Message)
	}
	if 0 != len(e.params.Data) {
		builder.Append(constant.Data).Append(constant.Equal).Append(e.params.Data.String())
	}
	builder.Append(constant.JsonEnd)

	return builder.String()
}
