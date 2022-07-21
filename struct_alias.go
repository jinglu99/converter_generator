package converter_generator

import "strconv"

var withPkgName = false
var structCounter = map[string]int64{}
var structName = map[string]string{}

var aliasMap = map[string]string{}

func addStructAlias(info typeInfo, alias string) {
	aliasMap[info.TypeString()] = alias
}

func getStructAlias(info typeInfo) string {
	alias, ok := aliasMap[info.TypeString()]
	if ok && len(alias) > 0 {
		return alias
	}

	if withPkgName {
		return toCamelCase(info.PkgName() + "_" + info.GetType().Name())
	}

	if name, ok := structName[info.TypeString()]; ok {
		return name
	}

	var name = info.GetType().Name()
	idx := structCounter[info.GetType().Name()]
	if idx > 0 {
		name = name + strconv.FormatInt(idx, 10)
	}
	structName[info.TypeString()] = name
	structCounter[info.GetType().Name()] = idx + 1
	return name
}
