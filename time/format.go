package time

import (
	"fmt"
	stdtime "time"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func Format() jsonnet.NativeFunction {
	return jsonnet.NativeFunction{
		Name:   "format",
		Params: ast.Identifiers{"epochMs", "layout"},
		Func: func(input []any) (any, error) {
			epochMs, ok := input[0].(float64)
			if !ok {
				return nil, fmt.Errorf("epochMs must be a number")
			}
			layout, ok := input[1].(string)
			if !ok {
				return nil, fmt.Errorf("layout must be a string")
			}
			return FormatTime(epochMs, layout), nil
		},
	}
}

func FormatTime(epochMs float64, layout string) string {
	return stdtime.UnixMilli(int64(epochMs)).Format(layout)
}
