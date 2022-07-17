package converter_generator

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"wangjinglu/autoconv/data"
)

func TestTypeInfoName(t *testing.T) {
	tp := typeInfo{}

	// Bool
	boolVal := true
	tp = typeInfo{t: reflect.TypeOf(boolVal)}
	assert.Equal(t, "Bool", tp.TypeName())

	// BoolPtr
	tp = typeInfo{t: reflect.TypeOf(&boolVal)}
	assert.Equal(t, "BoolPtr", tp.TypeName())

	// Int
	intVal := int(1)
	tp = typeInfo{t: reflect.TypeOf(intVal)}
	assert.Equal(t, "Int", tp.TypeName())

	// Int32
	int32Val := int32(1)
	tp = typeInfo{t: reflect.TypeOf(int32Val)}
	assert.Equal(t, "Int32", tp.TypeName())

	// IntList
	intListVal := make([]int, 0)
	tp = typeInfo{t: reflect.TypeOf(intListVal)}
	assert.Equal(t, "IntList", tp.TypeName())

	// IntPtrList
	intPtrList := make([]*int, 0)
	tp = typeInfo{t: reflect.TypeOf(intPtrList)}
	assert.Equal(t, "IntPtrList", tp.TypeName())

	// StringIntMap
	stringIntMapVal := map[string]int{}
	tp = typeInfo{t: reflect.TypeOf(stringIntMapVal)}
	assert.Equal(t, "StringIntMap", tp.TypeName())

	// StringIntMap
	stringIntPtrMapVal := map[string]*int{}
	tp = typeInfo{t: reflect.TypeOf(stringIntPtrMapVal)}
	assert.Equal(t, "StringIntPtrMap", tp.TypeName())

	// StringIntMapPtr
	tp = typeInfo{t: reflect.TypeOf(&stringIntPtrMapVal)}
	assert.Equal(t, "StringIntPtrMapPtr", tp.TypeName())

	// StringInterfaceMap
	stringInterfaceMap := map[string]interface{}{}
	tp = typeInfo{t: reflect.TypeOf(stringInterfaceMap)}
	assert.Equal(t, "StringInterfaceMap", tp.TypeName())

	// struct
	dataAVal := data.DataA{}
	tp = typeInfo{t: reflect.TypeOf(dataAVal)}
	assert.Equal(t, "DataA", tp.TypeName())

	// enum
	enumA := data.EnumA(1)
	tp = typeInfo{t: reflect.TypeOf(enumA)}
	assert.Equal(t, "EnumA", tp.TypeName())
}

func TestTypeInfoTypeString(t *testing.T) {
	tp := typeInfo{}

	// bool
	boolVal := true
	tp = typeInfo{t: reflect.TypeOf(boolVal)}
	assert.Equal(t, "bool", tp.TypeString())

	// *bool
	tp = typeInfo{t: reflect.TypeOf(&boolVal)}
	assert.Equal(t, "*bool", tp.TypeString())

	// int
	intVal := int(1)
	tp = typeInfo{t: reflect.TypeOf(intVal)}
	assert.Equal(t, "int", tp.TypeString())

	// int32
	int32Val := int32(1)
	tp = typeInfo{t: reflect.TypeOf(int32Val)}
	assert.Equal(t, "int32", tp.TypeString())

	// []int
	intListVal := make([]int, 0)
	tp = typeInfo{t: reflect.TypeOf(intListVal)}
	assert.Equal(t, "[]int", tp.TypeString())

	// []*int
	intPtrList := make([]*int, 0)
	tp = typeInfo{t: reflect.TypeOf(intPtrList)}
	assert.Equal(t, "[]*int", tp.TypeString())

	// map[string]int
	stringIntMapVal := map[string]int{}
	tp = typeInfo{t: reflect.TypeOf(stringIntMapVal)}
	assert.Equal(t, "map[string]int", tp.TypeString())

	// map[string]*int
	stringIntPtrMapVal := map[string]*int{}
	tp = typeInfo{t: reflect.TypeOf(stringIntPtrMapVal)}
	assert.Equal(t, "map[string]*int", tp.TypeString())

	// *map[string]*int
	tp = typeInfo{t: reflect.TypeOf(&stringIntPtrMapVal)}
	assert.Equal(t, "*map[string]*int", tp.TypeString())

	// map[string]interface{}
	stringInterfaceMap := map[string]interface{}{}
	tp = typeInfo{t: reflect.TypeOf(stringInterfaceMap)}
	assert.Equal(t, "map[string]interface{}", tp.TypeString())

	// struct
	dataAVal := data.DataA{}
	tp = typeInfo{t: reflect.TypeOf(&dataAVal)}
	assert.Equal(t, "*autoconv_data.DataA", tp.TypeString())

	// enum
	enumAVal := data.EnumA(1)
	tp = typeInfo{t: reflect.TypeOf(&enumAVal)}
	assert.Equal(t, "*autoconv_data.EnumA", tp.TypeString())
}
