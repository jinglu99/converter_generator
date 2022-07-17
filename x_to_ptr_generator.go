package converter_generator

import (
	"bytes"
	"html/template"
	"reflect"
)

func newX2PtrGenerator() generator {
	return x2PtrGenerator{}
}

type x2PtrGenerator struct {
}

const x2PtrTpl = `
{{- if .needCheckNil }}if in == nil { 
	return nil 
}
{{ end }}
{{- if .directlyAssign }} return &in
{{- else -}} 
out := {{.convFunc}}(in)
return &out
{{- end -}}
`

var x2PtrTplParser = template.Must(template.New("x2Ptr").Parse(x2PtrTpl))

func (x x2PtrGenerator) Handle(in, out typeInfo) string {
	var needCheckNil bool
	var directlyAssign bool
	var convFunc string
	var outType string

	if in.Kind() == reflect.Ptr {
		needCheckNil = true
	}

	trueType := out.Elem()
	if trueType.AssignableTo(out) {
		directlyAssign = true
	} else {
		convFunc = convertByTypeInfo(in, trueType).FuncName()
	}

	argvs := map[string]interface{}{
		"needCheckNil":   needCheckNil,
		"convFunc":       convFunc,
		"directlyAssign": directlyAssign,
		"outType":        outType,
	}

	buf := bytes.Buffer{}
	err := x2PtrTplParser.Execute(&buf, argvs)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
