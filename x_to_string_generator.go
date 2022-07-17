package converter_generator

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
)

func newX2StringGenerator() generator {
	return x2StringGenerator{}
}

type x2StringGenerator struct {
}

const x2StringTpl = `
{{- if .needCheckNil }}if in == nil { 
	return ""
}
{{- end }}
return {{.outType}}({{.trulyIn}})`

var x2StringTplParser = template.Must(template.New("x2Str").Parse(x2StringTpl))

func (x x2StringGenerator) Handle(in, out typeInfo) string {
	oriIn, oriOut := in, out

	var trulyIn string
	var outType string
	var needCheckNil bool

	if in.Kind() == reflect.Ptr {
		needCheckNil = true
	}

	trulyIn = "in"
	for true {
		if in.Kind() == reflect.Ptr {
			trulyIn = "*" + trulyIn
			in = in.Elem()
			continue
		}
		break
	}

	outType = out.TypeString()

	switch in.Kind() {
	case reflect.String:
	default:
		panic(fmt.Sprintf("can't convert from %v to %v", oriIn.TypeString(), oriOut.TypeString()))
	}

	argvs := map[string]interface{}{
		"needCheckNil": needCheckNil,
		"outType":      outType,
		"trulyIn":      trulyIn,
	}

	buf := bytes.Buffer{}
	err := x2StringTplParser.Execute(&buf, argvs)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
