package main

import (
	"github.com/jingleWang/converter_generator"
	"github.com/jingleWang/converter_generator/test/data"
)

func main() {
	cg := converter_generator.ConverterGenerator{}
	cg.OutputDir("converts")
	cg.PkgName("converts")
	//cg.FileName("xxx.go")

	//cg.RegisterCustomConverter(time.Time{}, int64(1), "ConvTimeToInt")
	//cg.RegisterCustomConverter(int64(1), time.Time{}, "ConvIntToTime")

	cg.Convert(data.D{}, data.C{}, converter_generator.ConversionConfig{
		FieldMapper: [][]string{
			{"StructA", "StructB"},
		},
		Export: true,
	})
	cg.Generate()

}
