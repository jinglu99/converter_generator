package converter_generator

var aliasMap = map[string]string{}

func addStructAlias(info typeInfo, alias string) {
	aliasMap[info.TypeString()] = alias
}

func getStructAlias(info typeInfo) string {
	alias, ok := aliasMap[info.TypeString()]
	if ok && len(alias) > 0 {
		return alias
	}

	return info.GetType().Name()
}
