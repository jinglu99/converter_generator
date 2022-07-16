package autoconv

import (
	"reflect"
)

var conversions = map[string]conversion{}

func convert(source interface{}, dest interface{}) {
	sType := newTypeInfo(reflect.TypeOf(source))
	dType := newTypeInfo(reflect.TypeOf(dest))
	convertByTypeInfo(sType, dType)
}

func convertByTypeInfo(sType typeInfo, dType typeInfo) conversion {
	if sType.GetType() == dType.GetType() {
		panic("input same struct")
	}

	key := sType.TypeString() + "_" + dType.TypeString()
	if c, ok := conversions[key]; ok {
		return c
	}

	c := conversion{
		sType: sType,
		dType: dType,
	}
	c.Generate()
	//fmt.Println(c.body)
	conversions[key] = c
	return c
}

