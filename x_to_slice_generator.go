package converter_generator

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
)

func newX2SliceGenerator() generator {
	return x2SliceGenerator{}
}

type x2SliceGenerator struct {
}

const x2SliceTpl = `if in == nil { 
	return nil 
}

out := make({{.outType}}, 0)
for _, item := range {{.trulyIn}} {
{{- if .directlyAssign }}
	out = append(out, item)
{{- else }}
	out = append(out, {{.convFunc}}(item))
{{- end }}
}
return out`

var x2SliceTplParser = template.Must(template.New("x2Slice").Parse(x2SliceTpl))

func (x x2SliceGenerator) Handle(in, out typeInfo) string {
	oriIn, oriOut := in, out

	var trulyIn string
	var directlyAssign bool
	var convFunc string
	var outType string

	trulyIn = "in"
	for true {
		if in.Kind() == reflect.Ptr {
			trulyIn = "*" + trulyIn
			in = in.Elem()
			continue
		}
		break
	}

	if in.Kind() != reflect.Slice && in.Kind() != reflect.Array {
		panic(fmt.Sprintf("can't convert from %v to %v", oriIn.TypeString(), oriOut.TypeString()))
	}

	outTrueType := out.Elem()
	inTrueType := in.Elem()
	if outTrueType.AssignableTo(inTrueType) {
		directlyAssign = true
	} else {
		convFunc = convertByTypeInfo(inTrueType, outTrueType).FuncName()
	}

	outType = oriOut.TypeString()

	argvs := map[string]interface{}{
		"outType":        outType,
		"convFunc":       convFunc,
		"directlyAssign": directlyAssign,
		"trulyIn":        trulyIn,
	}

	buf := bytes.Buffer{}
	err := x2SliceTplParser.Execute(&buf, argvs)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
