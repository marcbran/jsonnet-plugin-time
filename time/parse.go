package time

import (
	"fmt"
	stdtime "time"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func Parse() jsonnet.NativeFunction {
	return jsonnet.NativeFunction{
		Name:   "parse",
		Params: ast.Identifiers{"value", "layout"},
		Func: func(input []any) (any, error) {
			value, ok := input[0].(string)
			if !ok {
				return nil, fmt.Errorf("value must be a string")
			}
			layout, ok := input[1].(string)
			if !ok {
				return nil, fmt.Errorf("layout must be a string")
			}
			t, err := ParseTime(value, layout)
			if err != nil {
				return nil, err
			}
			return float64(t.UnixMilli()), nil
		},
	}
}

func ParseTime(value, layout string) (stdtime.Time, error) {
	t, err := stdtime.Parse(layout, value)
	if err != nil {
		return stdtime.Time{}, fmt.Errorf("invalid time %q for layout %q: %w", value, layout, err)
	}
	return t, nil
}
