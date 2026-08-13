package time

import (
	"fmt"
	"strconv"
	"strings"
	stdtime "time"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func AddDuration() jsonnet.NativeFunction {
	return jsonnet.NativeFunction{
		Name:   "addDuration",
		Params: ast.Identifiers{"epochMs", "spec"},
		Func: func(input []any) (any, error) {
			epochMs, ok := input[0].(float64)
			if !ok {
				return nil, fmt.Errorf("epochMs must be a number")
			}
			spec, ok := input[1].(string)
			if !ok {
				return nil, fmt.Errorf("spec must be a string")
			}
			result, err := AddDurationSpec(stdtime.UnixMilli(int64(epochMs)), spec)
			if err != nil {
				return nil, err
			}
			return float64(result.UnixMilli()), nil
		},
	}
}

func AddDurationSpec(ref stdtime.Time, spec string) (stdtime.Time, error) {
	trimmed := strings.TrimSpace(spec)
	if trimmed == "" {
		return stdtime.Time{}, fmt.Errorf("empty duration")
	}

	sign := 1
	rest := trimmed
	switch rest[0] {
	case '-':
		sign = -1
		rest = rest[1:]
	case '+':
		rest = rest[1:]
	}
	if rest == "" {
		return stdtime.Time{}, fmt.Errorf("invalid duration %q: no digits after sign", spec)
	}

	rest, years, months, days := consumeCalendarUnits(rest)
	result := ref.AddDate(sign*years, sign*months, sign*days)

	if rest != "" {
		d, err := stdtime.ParseDuration(rest)
		if err != nil {
			return stdtime.Time{}, fmt.Errorf("invalid duration %q: %w", spec, err)
		}
		result = result.Add(stdtime.Duration(sign) * d)
	}

	return result, nil
}

func consumeCalendarUnits(s string) (rest string, years, months, days int) {
	units := []struct {
		suffix byte
		apply  func(n int)
	}{
		{'y', func(n int) { years += n }},
		{'M', func(n int) { months += n }},
		{'w', func(n int) { days += n * 7 }},
		{'d', func(n int) { days += n }},
	}

	for _, unit := range units {
		i := 0
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			i++
		}
		if i == 0 || i >= len(s) || s[i] != unit.suffix {
			continue
		}
		n, err := strconv.Atoi(s[:i])
		if err != nil {
			continue
		}
		unit.apply(n)
		s = s[i+1:]
	}
	return s, years, months, days
}
