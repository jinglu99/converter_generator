package converter_generator

import (
	"bytes"
	"html/template"
	"reflect"
	"unicode"
)

func newX2StructGenerator() generator {
	return x2StructGenerator{}
}

type x2StructGenerator struct {
}

const x2structTpl = `
{{- if .needCheckNil }}if in == nil { 
	return {{ .outType}}{} 
}
{{ end }}
{{- if .directlyAssign }}
	return {{.trulyIn}}
{{- else -}}
	return {{ .outType}}{
{{ range .fields -}}
	{{ if .fieldDirectAssign -}}
	{{.name}}: in.{{.source}},
	{{- else -}}
	{{.name}}: {{.fieldConvFunc}}(in.{{.source}}),
	{{- end }}
{{ end -}}
} 
{{- end }}`

var x2StructTplParser = template.Must(template.New("x2struct").Parse(x2structTpl))

func (x x2StructGenerator) Handle(in, out typeInfo) string {
	_, oriOut := in, out

	var needCheckNil bool
	var trulyIn string
	var directlyAssign bool
	var convertible bool
	var convFunc string
	var outType string
	var fields []map[string]interface{}

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

	outTrueType := out
	inTrueType := in
	if outTrueType.Equal(inTrueType) {
		directlyAssign = true
	} else {
		fields = make([]map[string]interface{}, 0)
		fieldsCn := out.GetType().NumField()
		for i := 0; i < fieldsCn; i++ {
			var name string
			var source string
			var fieldDirectAssign bool
			var fieldConvFunc string

			field := outTrueType.GetType().Field(i)
			if unicode.IsLower([]rune(field.Name)[0]) {
				continue
			}

			sourceFiledName, customConvFunc := getConvertSourceFieldName(inTrueType, outTrueType, field.Name)
			sourceField, ok := inTrueType.GetType().FieldByName(sourceFiledName)
			if !ok {
				continue
			}

			name = field.Name
			source = sourceField.Name

			fieldTypeInfo := newTypeInfo(field.Type)
			sourceFieldInfo := newTypeInfo(sourceField.Type)
			if customConvFunc != nil && len(*customConvFunc) > 0 {
				fieldConvFunc = *customConvFunc
			} else if fieldTypeInfo.AssignableTo(sourceFieldInfo) {
				fieldDirectAssign = true
			} else {
				fieldConvFunc = convertByTypeInfo(sourceFieldInfo, fieldTypeInfo).FuncName()
			}

			fields = append(fields, map[string]interface{}{
				"name":              name,
				"source":            source,
				"fieldDirectAssign": fieldDirectAssign,
				"fieldConvFunc":     fieldConvFunc,
			})

		}
	}

	outType = oriOut.TypeString()

	argvs := map[string]interface{}{
		"outType":        outType,
		"convertible":    convertible,
		"convFunc":       convFunc,
		"directlyAssign": directlyAssign,
		"trulyIn":        trulyIn,
		"needCheckNil":   needCheckNil,
		"fields":         fields,
	}

	buf := bytes.Buffer{}
	err := x2StructTplParser.Execute(&buf, argvs)
	if err != nil {
		panic(err)
	}
	return buf.String()

}
