package autoconv

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
{{- if .convertible }} return &in 
{{- else -}} 
out := {{.convFunc}}(in)
return &out
{{- end -}}
`

var x2PtrTplParser = template.Must(template.New("x2Ptr").Parse(x2PtrTpl))

func (x x2PtrGenerator) Handle(in, out typeInfo) string {
	var needCheckNil bool
	var convertible bool
	var convFunc string

	if in.Kind() == reflect.Ptr {
		needCheckNil = true
	}

	trueType := out.Elem()
	if trueType.ConvertibleTo(in) {
		convertible = true
	} else {
		convFunc = convertByTypeInfo(in, trueType).FuncName()
	}

	argvs := map[string]interface{}{
		"needCheckNil": needCheckNil,
		"convertible":  convertible,
		"convFunc":     convFunc,
	}

	buf := bytes.Buffer{}
	err := x2PtrTplParser.Execute(&buf, argvs)
	if err != nil {
		panic(err)
	}
	return buf.String()
	//
	//sb := strings.Builder{}
	//if in.Kind() == reflect.Ptr {
	//	checkNil(&sb)
	//}
	//trueType := out.Elem()
	//if trueType.ConvertibleTo(in) {
	//	sb.WriteString("return &in")
	//} else {
	//	convFunc := convertByTypeInfo(in, trueType).FuncName()
	//	sb.WriteString(fmt.Sprintf("out := %v(in)\n", convFunc))
	//	sb.WriteString("return &out")
	//}
	//return sb.String()
}
