package converter_generator

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
)

func newX2BoolGenerator() generator {
	return x2BoolGenerator{}
}

type x2BoolGenerator struct {
}

const x2BoolTpl = `
{{- if .needCheckNil }}if in == nil { 
	return false 
}
{{- end }}
return {{.outType}}({{.trulyIn}})`

var x2BoolTplParser = template.Must(template.New("x2Bool").Parse(x2BoolTpl))

func (x x2BoolGenerator) Handle(in, out typeInfo) string {
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
	case reflect.Bool:
	default:
		panic(fmt.Sprintf("can't convert from %v to %v", oriIn.TypeString(), oriOut.TypeString()))
	}

	argvs := map[string]interface{}{
		"needCheckNil": needCheckNil,
		"outType":      outType,
		"trulyIn":      trulyIn,
	}

	buf := bytes.Buffer{}
	err := x2BoolTplParser.Execute(&buf, argvs)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
