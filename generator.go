package autoconv

type generator interface {
	Handle(in, out typeInfo) string
}
