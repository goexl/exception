package exception_test

import (
	"testing"

	"github.com/goexl/exception"
	"github.com/stretchr/testify/require"
)

func TestScheduling(t *testing.T) {
	_exception := exception.New().Build()
	require.NotNil(t, _exception, "创建异常出错")
}
