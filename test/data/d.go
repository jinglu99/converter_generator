package data

import "time"

type EnumA int64
type EnumB int64

type A struct {
	Str1             string
	Str2             *string
	Int              int
	Int1             *int
	StringStringMap1 map[string]string
	StringStringMap2 map[string]*string
	B                *bool
	T                time.Time
	t                time.Time
	EA               EnumA
}

type B struct {
	Str1             string
	Str2             string
	Int              int64
	Int1             *int64
	StringStringMap1 map[*string]string
	StringStringMap2 map[string]*string
	B                bool
	T                int64
	t                time.Time
	EA               EnumB
}

type C struct {
	StructB B
	Data    map[string]A
}

type D struct {
	StructA A
	Data    map[string]B
}
