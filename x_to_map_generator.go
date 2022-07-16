package autoconv

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
{{- else if .keyConvertible}}
	key := {{.outKeyType}}(k)
{{- else }}
	key := {{.convKeyFunc}}(k)
{{- end }}
{{- if .directlyAssignVal }}
	val := v
{{- else if .valConvertible}}
	val := {{.outValType}}(v)
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
	var keyConvertible bool
	var valConvertible bool
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

	if inTrueKeyType.Equal(outTrueKeyType) {
		directlyAssignKey = true
	} else if inTrueKeyType.ConvertibleTo(outTrueKeyType) {
		keyConvertible = true
	} else {
		convKeyFunc = convertByTypeInfo(inTrueKeyType, outTrueKeyType).FuncName()
	}

	if inTrueValType.Equal(outTrueValType) {
		directlyAssignVal = true
	} else if inTrueValType.ConvertibleTo(outTrueValType) {
		keyConvertible = true
	} else {
		convValFunc = convertByTypeInfo(inTrueValType, outTrueValType).FuncName()
	}

	outType = oriOut.TypeString()
	argvs := map[string]interface{}{
		"trulyIn":           trulyIn,
		"directlyAssignKey": directlyAssignKey,
		"directlyAssignVal": directlyAssignVal,
		"keyConvertible": keyConvertible,
		"valConvertible": valConvertible,
		"convKeyFunc":convKeyFunc,
		"convValFunc": convValFunc,
		"outKeyType": outKeyType,
		"outValType": outValType,
		"outType": outType,
	}

	buf := bytes.Buffer{}
	err := x2MapTplParser.Execute(&buf, argvs)
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
	//if in.Kind() != reflect.Map {
	//	panic(fmt.Sprintf("can't convert from %v to %v", oriIn.TypeString(), oriOut.TypeString()))
	//}
	//
	//sb.WriteString(fmt.Sprintf("for k, v := range %v {\n", inVal))
	//
	//outTrueKey := out.Key()
	//outTrueType := out.Elem()
	//inTrueKey := in.Key()
	//inTrueType := in.Elem()
	//
	//if inTrueKey.Equal(outTrueKey) {
	//	sb.WriteString("key := k\n")
	//} else if inTrueKey.ConvertibleTo(outTrueKey) {
	//	sb.WriteString(fmt.Sprintf("key := %v(k)\n", outTrueKey.TypeString()))
	//} else {
	//	sb.WriteString(fmt.Sprintf("key := %v(k)\n", convertByTypeInfo(inTrueKey, outTrueKey).FuncName()))
	//}
	//
	//if inTrueType.Equal(outTrueType) {
	//	sb.WriteString("val := k\n")
	//} else if inTrueType.ConvertibleTo(outTrueType) {
	//	sb.WriteString(fmt.Sprintf("val := %v(k)\n", outTrueType.TypeString()))
	//} else {
	//	sb.WriteString(fmt.Sprintf("val := %v(k)\n", convertByTypeInfo(inTrueType, outTrueType).FuncName()))
	//}
	//
	//sb.WriteString(fmt.Sprintf("out[key]=val\n"))
	//
	//sb.WriteString("}\n")
	//sb.WriteString("return out")
	//return sb.String()
}
