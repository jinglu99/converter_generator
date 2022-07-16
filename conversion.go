package autoconv

import (
	"fmt"
	"reflect"
	"strings"
)

type conversion struct {
	sType, dType typeInfo
	body         string
}

func getAllPkgs() map[string]string {
	pkgs := map[string]string{}
	for _, conversion := range conversions {
		var pkg, path string
		pkg = conversion.dType.PkgName()
		path = conversion.dType.PkgPath()
		if pkg != "" || path != "" {
			pkgs[pkg] = path
		}

		pkg = conversion.sType.PkgName()
		path = conversion.sType.PkgPath()
		if pkg != "" || path != "" {
			pkgs[pkg] = path
		}
	}
	return pkgs
}

func (c *conversion) Generate() {
	funcName := c.FuncName()

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("func %v(in %v) %v", funcName, c.sType.TypeString(), c.dType.TypeString()))
	sb.WriteString(" {\n")
	sb.WriteString(c.generateBody())
	sb.WriteString("\n}")
	c.body = sb.String()
}

func (c conversion) generateBody() string {
	sb := strings.Builder{}

	if c.sType.ConvertibleTo(c.dType) {
		sb.WriteString(fmt.Sprintf("return %v(in)", c.dType.TypeString()))
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
	default:
		panic(fmt.Sprintf("can't convert from %v to %v", in.TypeString(), out.TypeString()))
	}

}

func (c conversion) FuncName() string {
	return fmt.Sprintf("Convert%vTo%v", c.sType.TypeName(), c.dType.TypeName())
}
