package data

type A struct {
	Str1              string
	Str2              *string
	Int               int
	Int1              *int
	StringStringMap1  map[string]string
	StringStringMap2 map[string]*string
}

type B struct {
	Str1              string
	Str2              string
	Int               int64
	Int1              *int64
	StringStringMap1  map[*string]string
	StringStringMap2 map[string]*string
}

type C struct {
	StructB B
	Data    map[string]A
}

type D struct {
	StructA A
	Data    map[string]B
}
