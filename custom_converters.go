package converter_generator

var _ conversion = (*customConversion)(nil)

type customConversion struct {
	sType, dType typeInfo
	f            string
}

func (c *customConversion) Body() string {
	panic("implement me")
}

func (c *customConversion) SType() typeInfo {
	return c.sType
}

func (c *customConversion) DType() typeInfo {
	return c.dType
}

func (c *customConversion) Generate() {
	panic("implement me")
}

func (c customConversion) FuncName() string {
	return c.f
}

var customFunc = map[string]conversion{}

func addCustomFunc(sType, dType typeInfo, name string) {
	key := generateConversionKey(sType, dType)
	customFunc[key] = &customConversion{
		sType: sType,
		dType: dType,
		f:     name,
	}
}

func getCustomFunc(sType, dType typeInfo) conversion {
	return customFunc[generateConversionKey(sType, dType)]
}
