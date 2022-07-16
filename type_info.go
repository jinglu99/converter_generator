package autoconv

import (
	"fmt"
	"reflect"
	"strings"
)

func newTypeInfo(t reflect.Type) typeInfo {
	return typeInfo{t: t}
}

type typeInfo struct {
	t reflect.Type
}

func (ti typeInfo) GetType() reflect.Type {
	return ti.t
}

func (ti typeInfo) PkgPath() string {
	return ti.t.PkgPath()
}

func (ti typeInfo) PkgName() string {
	pkgLen := 2

	path := ti.PkgPath()
	if path == "" {
		return ""
	}

	pkgs := strings.Split(path, "/")
	pkgs = pkgs[1:]
	if len(pkgs) <= pkgLen {
		return strings.ReplaceAll(strings.Join(pkgs, "_"), ".", "_")
	}

	return strings.ReplaceAll(strings.Join(pkgs[len(pkgs)-pkgLen:], "_"), ".", "_")
}

func (ti typeInfo) TypeName() string {
	t := ti.t
	switch t.Kind() {
	case reflect.Bool, reflect.Int,
		reflect.Int8, reflect.Int32,
		reflect.Int64, reflect.Uint,
		reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Float32,
		reflect.Float64, reflect.String:
		return strings.ToUpper(t.Name()[:1]) + t.Name()[1:]
	case reflect.Array, reflect.Slice:
		tmpInfo := typeInfo{t.Elem()}
		return tmpInfo.TypeName() + "List"
	case reflect.Map:
		kInfo := typeInfo{t.Key()}
		vInfo := typeInfo{t.Elem()}
		return kInfo.TypeName() + vInfo.TypeName() + "Map"
	case reflect.Struct:
		return t.Name()
	case reflect.Ptr:
		tmpTI := typeInfo{t: t.Elem()}
		return tmpTI.TypeName() + "Ptr"
	case reflect.Interface:
		return "Interface"
	default:
		panic(fmt.Sprintf("not support type_info: %v", t))
	}
}

func (ti typeInfo) TypeString() string {
	t := ti.t
	switch t.Kind() {
	case reflect.Bool, reflect.Int,
		reflect.Int8, reflect.Int32,
		reflect.Int64, reflect.Uint,
		reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Float32,
		reflect.Float64, reflect.String:
		pkgName := ti.PkgName()
		if pkgName == "" {
			return t.Name()
		}
		return pkgName + "." + t.Name()
	case reflect.Array, reflect.Slice:
		tmpInfo := typeInfo{t.Elem()}
		return "[]" + tmpInfo.TypeString()
	case reflect.Map:
		kInfo := typeInfo{t.Key()}
		vInfo := typeInfo{t.Elem()}
		return "map[" + kInfo.TypeString() + "]" + vInfo.TypeString()
	case reflect.Struct:
		pkgName := ti.PkgName()
		if pkgName == "" {
			return t.Name()
		}
		return pkgName + "." + t.Name()
	case reflect.Ptr:
		tmpTI := typeInfo{t: t.Elem()}
		return "*" + tmpTI.TypeString()
	case reflect.Interface:
		return "interface{}"

	default:
		panic(fmt.Sprintf("not support type_info: %v", t))
	}
}

func (ti typeInfo) Equal(t typeInfo) bool {
	return t.t == ti.t && t.TypeName() == ti.TypeName()
}

func (ti typeInfo) ConvertibleTo(t typeInfo) bool {
	return ti.t.ConvertibleTo(t.t)
}

func (ti typeInfo) Kind() reflect.Kind {
	return ti.t.Kind()
}

func (ti typeInfo) Elem() typeInfo {
	return newTypeInfo(ti.t.Elem())
}

func (ti typeInfo) Key() typeInfo {
	return newTypeInfo(ti.t.Key())
}
