package converts

import (
	test_data "github.com/jingleWang/converter_generator/test/data"
	test_data2 "github.com/jingleWang/converter_generator/test/data2"
	test_data3 "github.com/jingleWang/converter_generator/test/data3"
)

func convertInt64PtrToInt(in *int64) int {
	if in == nil {
		return 0
	}
	return int(*in)
}

func convertIntToInt64(in int) int64 {

	return int64(in)
}

func convertIntPtrToInt64(in *int) int64 {
	if in == nil {
		return 0
	}
	return int64(*in)
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

func convertBToA(in test_data.B) test_data.A {
	return test_data.A{
		Str1:             in.Str1,
		Str2:             convertStringToStringPtr(in.Str2),
		Int:              convertInt64ToInt(in.Int),
		Int1:             convertInt64PtrToIntPtr(in.Int1),
		StringStringMap1: convertStringPtrStringMapToStringStringMap(in.StringStringMap1),
		StringStringMap2: in.StringStringMap2,
		B:                convertBoolToBoolPtr(in.B),
		T:                ConvIntToTime(in.T),
		t:                in.t,
	}
}

func convertIntPtrToInt64Ptr(in *int) *int64 {
	if in == nil {
		return nil
	}
	out := convertIntPtrToInt64(in)
	return &out
}

func convertA1ToA2(in test_data2.A) test_data3.A {
	return test_data3.A{
		Str1:             in.Str1,
		Str2:             in.Str2,
		Int:              in.Int,
		Int1:             in.Int1,
		StringStringMap1: in.StringStringMap1,
		StringStringMap2: in.StringStringMap2,
		B:                in.B,
		T:                in.T,
	}
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

func convertBoolToBoolPtr(in bool) *bool {
	return &in
}

func convertInt64ToInt(in int64) int {

	return int(in)
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
		T:                ConvTimeToInt(in.T),
		t:                in.t,
	}
}

func convertStringToStringPtr(in string) *string {
	return &in
}

func convertBoolPtrToBool(in *bool) bool {
	if in == nil {
		return false
	}
	return bool(*in)
}

func ConvertDToC(in test_data.D) test_data.C {
	return test_data.C{
		Data: convertStringBMapToStringAMap(in.Data),
	}
}

func convertStringPtrToString(in *string) string {
	if in == nil {
		return ""
	}
	return string(*in)
}

func convertAToA1(in test_data.A) test_data2.A {
	return test_data2.A{
		Str1:             in.Str1,
		Str2:             in.Str2,
		Int:              in.Int,
		Int1:             in.Int1,
		StringStringMap1: in.StringStringMap1,
		StringStringMap2: in.StringStringMap2,
		B:                in.B,
		T:                in.T,
	}
}

func convertInt64PtrToIntPtr(in *int64) *int {
	if in == nil {
		return nil
	}
	out := convertInt64PtrToInt(in)
	return &out
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
