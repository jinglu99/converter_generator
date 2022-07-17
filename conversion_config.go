package converter_generator

var conversionConfigs = map[string]ConversionConfig{}

type ConversionConfig struct {
	sType, dType typeInfo
	key          string
	FieldMapper  [][]string
	fieldDict    map[string]string
	Name         string
	Export       bool
}

func (cc *ConversionConfig) setTypes(sType, dType typeInfo) {
	cc.sType = sType
	cc.dType = dType
	cc.key = generateConversionKey(sType, dType)
}

func addConversionConfig(sType, dType typeInfo, config ConversionConfig) {
	config.setTypes(sType, dType)
	conversionConfigs[generateConversionKey(sType, dType)] = config
}

func getConversionConfig(sType, dType typeInfo) *ConversionConfig {
	config, ok := conversionConfigs[generateConversionKey(sType, dType)]
	if !ok {
		return nil
	}

	return &config
}

func getConvertSourceFieldName(sType, dType typeInfo, field string) string {
	cc := getConversionConfig(sType, dType)
	if cc == nil {
		return field
	}

	if cc.fieldDict == nil {
		cc.fieldDict = map[string]string{}
		for _, mapper := range cc.FieldMapper {
			cc.fieldDict[mapper[1]] = mapper[0]
		}
	}

	source, ok := cc.fieldDict[field]
	if ok {
		return source
	}
	return field
}
