package main

import (
	"github.com/jingleWang/converter_generator"
	"github.com/jingleWang/converter_generator/test/data"
	"github.com/jingleWang/converter_generator/test/data2"
	"github.com/jingleWang/converter_generator/test/data3"
	"time"
)

func main() {
	cg := converter_generator.ConverterGenerator{}
	cg.OutputDir("converts")
	cg.PkgName("converts")
	//cg.FileName("xxx.go")

	cg.RegisterCustomConverter(time.Time{}, int64(1), "ConvTimeToInt")
	cg.RegisterCustomConverter(int64(1), time.Time{}, "ConvIntToTime")

	cg.Convert(data.A{}, data2.A{})
	cg.Convert(data2.A{}, data3.A{})

	cg.Convert(data.D{}, data.C{}, converter_generator.ConversionConfig{
		FieldMapper: []*converter_generator.FieldMapper{
			{"StructA", "StructB", nil},
		},
		Export: true,
	})
	cg.Generate()

}
