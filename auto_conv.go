package converter_generator

import (
	"reflect"
)

var conversions = map[string]conversion{}

func convert(source interface{}, dest interface{}) {
	sType := newTypeInfo(reflect.TypeOf(source))
	dType := newTypeInfo(reflect.TypeOf(dest))
	convertByTypeInfo(sType, dType)
}

func getAllPkgs() map[string]string {
	pkgs := map[string]string{}
	for _, conversion := range conversions {
		var pkg, path string
		pkg = conversion.SType().PkgName()
		path = conversion.SType().PkgPath()
		if pkg != "" || path != "" {
			pkgs[pkg] = path
		}

		pkg = conversion.DType().PkgName()
		path = conversion.DType().PkgPath()
		if pkg != "" || path != "" {
			pkgs[pkg] = path
		}
	}
	return pkgs
}

func convertByTypeInfo(sType typeInfo, dType typeInfo) conversion {
	if sType.GetType() == dType.GetType() {
		panic("input same struct: " + sType.TypeString())
	}

	key := generateConversionKey(sType, dType)
	if c, ok := customFunc[key]; ok {
		return c
	}

	if c, ok := conversions[key]; ok {
		return c
	}

	c := &defaultConversion{
		sType: sType,
		dType: dType,
	}
	c.Generate()
	//fmt.Println(c.body)
	conversions[key] = c
	return c
}

func generateConversionKey(sType, dType typeInfo) string {
	return sType.TypeString() + "_" + dType.TypeString()
}
