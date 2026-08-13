package time

import (
	"github.com/google/go-jsonnet"
	"github.com/marcbran/jpoet/pkg/jpoet"
)

func Plugin(name string) *jpoet.Plugin {
	return jpoet.NewPlugin(name, []jsonnet.NativeFunction{
		Now(),
		AddDuration(),
		ParseRFC3339(),
	})
}
