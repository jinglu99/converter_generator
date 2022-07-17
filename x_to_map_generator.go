package converter_generator

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
)

func newX2MapGenerator() generator {
	return x2MapGenerator{}
}

type x2MapGenerator struct {
}

const x2MapTpl = `if in == nil { 
	return nil 
}

out := {{.outType}}{}
for k, v := range {{.trulyIn}} {
{{- if .directlyAssignKey }}
	key := k
{{- else }}
	key := {{.convKeyFunc}}(k)
{{- end }}
{{- if .directlyAssignVal }}
	val := v
{{- else }}
	val := {{.convValFunc}}(v)
{{- end }}
	out[key]=val
}
return out`

var x2MapTplParser = template.Must(template.New("x2Slice").Parse(x2MapTpl))

func (x x2MapGenerator) Handle(in, out typeInfo) string {
	oriIn, oriOut := in, out

	var trulyIn string
	var directlyAssignKey bool
	var directlyAssignVal bool
	var convKeyFunc string
	var convValFunc string

	var outType string
	var outKeyType string
	var outValType string

	trulyIn = "in"
	for true {
		if in.Kind() == reflect.Ptr {
			trulyIn = "*" + trulyIn
			in = in.Elem()
			continue
		}
		break
	}

	if in.Kind() != reflect.Map {
		panic(fmt.Sprintf("can't convert from %v to %v", oriIn.TypeString(), oriOut.TypeString()))
	}

	outTrueKeyType := out.Key()
	outTrueValType := out.Elem()
	inTrueKeyType := in.Key()
	inTrueValType := in.Elem()

	if inTrueKeyType.AssignableTo(outTrueKeyType) {
		directlyAssignKey = true
	} else {
		convKeyFunc = convertByTypeInfo(inTrueKeyType, outTrueKeyType).FuncName()
	}

	if inTrueValType.AssignableTo(outTrueValType) {
		directlyAssignVal = true
	} else {
		convValFunc = convertByTypeInfo(inTrueValType, outTrueValType).FuncName()
	}

	outType = oriOut.TypeString()
	argvs := map[string]interface{}{
		"trulyIn":           trulyIn,
		"directlyAssignKey": directlyAssignKey,
		"directlyAssignVal": directlyAssignVal,
		"convKeyFunc":       convKeyFunc,
		"convValFunc":       convValFunc,
		"outKeyType":        outKeyType,
		"outValType":        outValType,
		"outType":           outType,
	}

	buf := bytes.Buffer{}
	err := x2MapTplParser.Execute(&buf, argvs)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
