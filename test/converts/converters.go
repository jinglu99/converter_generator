package converts

import (
	test_data "github.com/jingleWang/converter_generator/test/data"
)

func convertIntToInt64(in int) int64 {

	return int64(in)
}

func convertStringToStringPtr(in string) *string {
	return &in
}

func convertStringPtrStringMapToStringStringMap(in map[*string]string) map[string]string {
	if in == nil {
		return nil
	}

	out := map[string]string{}
	for k, v := range in {
		key := convertStringPtrToString(k)
		val := v
		out[key] = val
	}
	return out
}

func convertStringBMapToStringAMap(in map[string]test_data.B) map[string]test_data.A {
	if in == nil {
		return nil
	}

	out := map[string]test_data.A{}
	for k, v := range in {
		key := k
		val := convertBToA(v)
		out[key] = val
	}
	return out
}

func convertInt64PtrToInt(in *int64) int {
	if in == nil {
		return 0
	}
	return int(*in)
}

func convertStringPtrToString(in *string) string {
	if in == nil {
		return ""
	}
	return string(*in)
}

func convertIntPtrToInt64(in *int) int64 {
	if in == nil {
		return 0
	}
	return int64(*in)
}

func convertAToB(in test_data.A) test_data.B {
	return test_data.B{
		Str1:             in.Str1,
		Str2:             convertStringPtrToString(in.Str2),
		Int:              convertIntToInt64(in.Int),
		Int1:             convertIntPtrToInt64Ptr(in.Int1),
		StringStringMap1: convertStringStringMapToStringPtrStringMap(in.StringStringMap1),
		StringStringMap2: in.StringStringMap2,
		B:                convertBoolPtrToBool(in.B),
	}
}

func convertInt64ToInt(in int64) int {

	return int(in)
}

func convertBoolPtrToBool(in *bool) bool {
	if in == nil {
		return false
	}
	return bool(*in)
}

func convertBoolToBoolPtr(in bool) *bool {
	return &in
}

func convertBToA(in test_data.B) test_data.A {
	return test_data.A{
		Str1:             in.Str1,
		Str2:             convertStringToStringPtr(in.Str2),
		Int:              convertInt64ToInt(in.Int),
		Int1:             convertInt64PtrToIntPtr(in.Int1),
		StringStringMap1: convertStringPtrStringMapToStringStringMap(in.StringStringMap1),
		StringStringMap2: in.StringStringMap2,
		B:                convertBoolToBoolPtr(in.B),
	}
}

func convertIntPtrToInt64Ptr(in *int) *int64 {
	if in == nil {
		return nil
	}
	out := convertIntPtrToInt64(in)
	return &out
}

func convertStringStringMapToStringPtrStringMap(in map[string]string) map[*string]string {
	if in == nil {
		return nil
	}

	out := map[*string]string{}
	for k, v := range in {
		key := convertStringToStringPtr(k)
		val := v
		out[key] = val
	}
	return out
}

func convertInt64PtrToIntPtr(in *int64) *int {
	if in == nil {
		return nil
	}
	out := convertInt64PtrToInt(in)
	return &out
}

func ConvertDToC(in test_data.D) test_data.C {
	return test_data.C{
		StructB: convertAToB(in.StructA),
		Data:    convertStringBMapToStringAMap(in.Data),
	}
}
