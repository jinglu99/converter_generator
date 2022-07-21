package data3

import "time"

type A struct {
	Str1             string
	Str2             *string
	Int              int
	Int1             *int
	StringStringMap1 map[string]string
	StringStringMap2 map[string]*string
	B                *bool
	T                time.Time
}
