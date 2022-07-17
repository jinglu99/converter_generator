package converter_generator

import (
	"io/ioutil"
	"reflect"
	"strings"
)

type ConverterGenerator struct {
	outputDir *string
	pkgName   *string
	fileName  *string

	conversions [][]typeInfo
}

func (cg *ConverterGenerator) OutputDir(dir string) *ConverterGenerator {
	cg.outputDir = &dir
	return cg
}

func (cg *ConverterGenerator) PkgName(pkg string) *ConverterGenerator {
	cg.pkgName = &pkg
	return cg
}


func (cg *ConverterGenerator) FileName(file string) *ConverterGenerator {
	cg.fileName = &file
	return cg
}

func (cg *ConverterGenerator) Alias(entity interface{}, alias string) {
	if entity == nil {
		panic("input struct can't be nil")
	}
	addStructAlias(newTypeInfo(reflect.TypeOf(entity)), alias)
}

func (cg *ConverterGenerator) Convert(a, b interface{}, config ...ConversionConfig) {
	if a == nil || b == nil {
		panic("struct can't be nil")
	}
	sType := newTypeInfo(reflect.TypeOf(a))
	dType := newTypeInfo(reflect.TypeOf(b))
	if len(config) > 0 {
		config[0].Export = true
		addConversionConfig(sType, dType, config[0])
	}

	cg.conversions = append(cg.conversions, []typeInfo{sType, dType})
}

func (cg *ConverterGenerator) Generate() {
	if cg.outputDir == nil || len(*cg.outputDir) == 0 {
		panic("must provide output dir")
	}

	if cg.pkgName == nil || len(*cg.pkgName) == 0 {
		panic("must provide pkg name")
	}

	for _, types := range cg.conversions {
		convertByTypeInfo(types[0], types[1])
	}

	cg.save()
}

func (cg ConverterGenerator) save() {
	fileName := "converters.go"
	if cg.fileName != nil && len(*cg.fileName) > 0 {
		fileName = *cg.fileName
	}

	output := strings.Builder{}
	output.WriteString("package " + *cg.pkgName + "\n\n")

	pkgs := getAllPkgs()
	if len(pkgs) > 0 {
		output.WriteString("import (\n")
		for k, v := range pkgs {
			output.WriteString(k + " \"" + v + "\"\n")
		}
		output.WriteString(")")
	}

	for _, conv := range conversions {
		output.WriteString("\n\n")
		output.WriteString(conv.body)
	}

	var filename = *cg.outputDir + "/" + fileName
	err := ioutil.WriteFile(filename, []byte(output.String()), 755)
	if err != nil {
		panic(err)
	}

}
