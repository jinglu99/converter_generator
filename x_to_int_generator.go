package converter_generator

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
)

func newX2IntGenerator() generator {
	return x2IntGenerator{}
}

type x2IntGenerator struct {
}

const x2IntTpl = `
{{- if .needCheckNil }}if in == nil { 
	return 0 
}
{{- end }}
return {{.outType}}({{.trulyIn}})
`

var x2IntTplParser = template.Must(template.New("x2Int").Parse(x2IntTpl))

func (x x2IntGenerator) Handle(in, out typeInfo) string {
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
	case reflect.Int,
		reflect.Int8, reflect.Int32,
		reflect.Int64, reflect.Uint,
		reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
	default:
		panic(fmt.Sprintf("can't convert from %v to %v", oriIn.TypeString(), oriOut.TypeString()))
	}

	argvs := map[string]interface{}{
		"needCheckNil": needCheckNil,
		"outType":      outType,
		"trulyIn":      trulyIn,
	}

	buf := bytes.Buffer{}
	err := x2IntTplParser.Execute(&buf, argvs)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
