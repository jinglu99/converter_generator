package data

type A struct {
	A interface{}
}

type B struct {
	B interface{}
}

type C struct {
	Data map[string]A
}

type D struct {
	Data map[string]B
}
