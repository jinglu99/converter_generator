package converter_generator

type generator interface {
	Handle(in, out typeInfo) string
}
