package converter_generator

import (
	"io/ioutil"
	"os"
	"strings"
)

type ConverterGenerator struct {
	OutputDir string
	PkgName   string
}

func (cg ConverterGenerator) Convert(a, b interface{}) {
	convert(a, b)
}

func (cg ConverterGenerator) Export() {

	output := strings.Builder{}
	output.WriteString("package " + cg.PkgName + "\n\n")

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

	//fmt.Println(output.String())

	os.MkdirAll(cg.OutputDir, 777)
	os.Chmod(cg.OutputDir, 0766)
	var filename = cg.OutputDir + "/convert.go"
	err := ioutil.WriteFile(filename, []byte(output.String()), 0766)
	if err != nil {
		panic(err)
	}

}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
