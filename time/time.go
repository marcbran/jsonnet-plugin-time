package time

import (
	stdtime "time"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func Now() jsonnet.NativeFunction {
	return jsonnet.NativeFunction{
		Name:   "now",
		Params: ast.Identifiers{},
		Func: func(input []any) (any, error) {
			return float64(stdtime.Now().UnixMilli()), nil
		},
	}
}
