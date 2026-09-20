package time

import (
	"github.com/google/go-jsonnet"
	"github.com/marcbran/jpoet/pkg/jpoet"
)

func Plugin(opts ...jpoet.PluginOption) *jpoet.Plugin {
	return jpoet.NewPlugin("time", []jsonnet.NativeFunction{
		Now(),
		AddDuration(),
		Parse(),
		Format(),
	}, opts...)
}
