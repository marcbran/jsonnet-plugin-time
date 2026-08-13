package time

import (
	"fmt"
	stdtime "time"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func ParseRFC3339() jsonnet.NativeFunction {
	return jsonnet.NativeFunction{
		Name:   "parseRFC3339",
		Params: ast.Identifiers{"value"},
		Func: func(input []any) (any, error) {
			value, ok := input[0].(string)
			if !ok {
				return nil, fmt.Errorf("value must be a string")
			}
			t, err := ParseRFC3339Time(value)
			if err != nil {
				return nil, err
			}
			return float64(t.UnixMilli()), nil
		},
	}
}

func ParseRFC3339Time(value string) (stdtime.Time, error) {
	t, err := stdtime.Parse(stdtime.RFC3339Nano, value)
	if err != nil {
		return stdtime.Time{}, fmt.Errorf("invalid RFC3339 time %q: %w", value, err)
	}
	return t, nil
}
