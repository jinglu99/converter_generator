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
{{- else if .convertible}}
	out = append(out, {{.outType}}(item))
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
	var convertible bool
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
	if outTrueType.Equal(inTrueType) {
		directlyAssign = true
	} else if inTrueType.ConvertibleTo(outTrueType) {
		convertible = true
	} else {
		convFunc = convertByTypeInfo(inTrueType, outTrueType).FuncName()
	}

	outType = oriOut.TypeString()

	argvs := map[string]interface{}{
		"outType":        outType,
		"convertible":    convertible,
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

	//sb := strings.Builder{}
	//checkNil(&sb)
	//sb.WriteString(fmt.Sprintf("out := make(%v, 0)\n", out.TypeString()))
	//inVal := "in"
	//for true {
	//	if in.Kind() == reflect.Ptr {
	//		inVal = "*" + inVal
	//		in = in.Elem()
	//		continue
	//	}
	//	break
	//}
	//if in.Kind() != reflect.Slice {
	//	panic(fmt.Sprintf("can't convert from %v to %v", oriIn.TypeString(), oriOut.TypeString()))
	//}
	//
	//sb.WriteString(fmt.Sprintf("for _, item := range %v {\n", inVal))
	//
	//outTrueType := out.Elem()
	//inTrueType := in.Elem()
	//if outTrueType.Equal(inTrueType) {
	//	sb.WriteString(fmt.Sprintf("out = append(out, item)"))
	//} else if inTrueType.ConvertibleTo(outTrueType) {
	//	sb.WriteString(fmt.Sprintf("out = append(out, %v(item))", outTrueType.TypeString()))
	//} else {
	//	convFunc := convertByTypeInfo(inTrueType, outTrueType).FuncName()
	//	sb.WriteString(fmt.Sprintf("out = append(out, %v(item))\n", convFunc))
	//}
	//sb.WriteString("\n}\n")
	//sb.WriteString("return out")
	//return sb.String()
}
