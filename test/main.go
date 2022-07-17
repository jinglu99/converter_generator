package main

import (
	"github.com/jingleWang/converter_generator"
	"github.com/jingleWang/converter_generator/test/data"
)

func main() {
	cg := converter_generator.ConverterGenerator{}
	cg.OutputDir("converts")
	cg.PkgName("converts")
	cg.FileName("xxx.go")

	cg.Convert(data.D{}, data.C{}, converter_generator.ConversionConfig{
		Export: true,
	})
	cg.Generate()
}
