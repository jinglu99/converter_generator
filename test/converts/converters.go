package converts

import (
	data "github.com/jingleWang/converter_generator/test/data"
)

func convertStringBMapToStringAMap(in map[string]data.B) map[string]data.A {
	if in == nil {
		return nil
	}

	out := map[string]data.A{}
	for k, v := range in {
		key := k
		val := convertDataBToDataA(v)
		out[key] = val
	}
	return out
}

func convertStringToStringPtr(in string) *string {
	return &in
}

func convertInt64ToInt(in int64) int {

	return int(in)
}

func convertBoolToBoolPtr(in bool) *bool {
	return &in
}

func ConvertDataDToDataC(in data.D) data.C {
	return data.C{
		Data: convertStringBMapToStringAMap(in.Data),
	}
}

func convertInt64PtrToIntPtr(in *int64) *int {
	if in == nil {
		return nil
	}
	out := convertInt64PtrToInt(in)
	return &out
}

func convertInt64PtrToInt(in *int64) int {
	if in == nil {
		return 0
	}
	return int(*in)
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

func convertStringPtrToString(in *string) string {
	if in == nil {
		return ""
	}
	return string(*in)
}

func convertDataBToDataA(in data.B) data.A {
	return data.A{
		Str1:             in.Str1,
		Str2:             convertStringToStringPtr(in.Str2),
		Int:              convertInt64ToInt(in.Int),
		Int1:             convertInt64PtrToIntPtr(in.Int1),
		StringStringMap1: convertStringPtrStringMapToStringStringMap(in.StringStringMap1),
		StringStringMap2: in.StringStringMap2,
		B:                convertBoolToBoolPtr(in.B),
		T:                ConvIntToTime(in.T),
	}
}
