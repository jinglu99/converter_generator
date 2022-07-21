package converter_generator

import (
	"fmt"
	"reflect"
	"strings"
)

type conversion interface {
	SType() typeInfo
	DType() typeInfo
	Generate()
	FuncName() string
	Body() string
}

var _ conversion = (*defaultConversion)(nil)

type defaultConversion struct {
	sType, dType typeInfo
	body         string
}

func (c *defaultConversion) Body() string {
	return c.body
}

func (c *defaultConversion) SType() typeInfo {
	return c.sType
}

func (c *defaultConversion) DType() typeInfo {
	return c.dType
}

func (c *defaultConversion) Generate() {
	funcName := c.FuncName()

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("func %v(in %v) %v", funcName, c.sType.TypeString(), c.dType.TypeString()))
	sb.WriteString(" {\n")
	sb.WriteString(c.generateBody())
	sb.WriteString("\n}")
	c.body = sb.String()
}

func (c defaultConversion) generateBody() string {
	sb := strings.Builder{}

	if c.sType.AssignableTo(c.dType) {
		sb.WriteString(fmt.Sprintf("return (%v)(in)", c.dType.TypeString()))
		return sb.String()
	}

	out := c.dType
	in := c.sType

	switch out.Kind() {
	case reflect.Ptr:
		return newX2PtrGenerator().Handle(in, out)
	case reflect.Slice:
		return newX2SliceGenerator().Handle(in, out)
	case reflect.Map:
		return newX2MapGenerator().Handle(in, out)
	case reflect.Struct:
		return newX2StructGenerator().Handle(in, out)
	case reflect.Int,
		reflect.Int8, reflect.Int32,
		reflect.Int64, reflect.Uint,
		reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return newX2IntGenerator().Handle(in, out)
	case reflect.String:
		return newX2StringGenerator().Handle(in, out)
	case reflect.Bool:
		return newX2BoolGenerator().Handle(in, out)
	default:
		fmt.Println(out.Kind())
		panic(fmt.Sprintf("can't convert from %v to %v", in.TypeString(), out.TypeString()))
	}

	return ""
}

func (c defaultConversion) FuncName() string {
	funcName := fmt.Sprintf("convert%vTo%v", c.sType.TypeName(), c.dType.TypeName())
	config := getConversionConfig(c.sType, c.dType)
	if config == nil {
		return funcName
	}

	if len(config.Name) > 0 {
		return config.Name
	}

	if config.Export {
		return strings.ToUpper(funcName[0:1]) + funcName[1:]
	}
	return funcName
}
