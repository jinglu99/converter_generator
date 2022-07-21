package converter_generator

var conversionConfigs = map[string]ConversionConfig{}

type FieldMapper struct {
	Source   string
	Target   string
	ConvFunc *string
}

type ConversionConfig struct {
	sType, dType typeInfo
	key          string
	FieldMapper  []*FieldMapper
	fieldDict    map[string]*FieldMapper
	Name         string
	Slim         bool
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

func getConvertSourceFieldName(sType, dType typeInfo, field string) (string, *string) {
	cc := getConversionConfig(sType, dType)
	if cc == nil {
		return field, nil
	}

	if cc.fieldDict == nil {
		cc.fieldDict = map[string]*FieldMapper{}
		for _, mapper := range cc.FieldMapper {
			//var convFunc *string = nil
			//if len(mapper) > 2 {
			//	convFunc = &mapper[2]
			//}
			cc.fieldDict[mapper.Source] = mapper
		}
	}

	source, ok := cc.fieldDict[field]
	if ok {
		return source.Target, source.ConvFunc
	}
	return field, nil
}
