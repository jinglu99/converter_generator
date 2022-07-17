package main

import (
	"github.com/jingleWang/converter_generator"
	"github.com/jingleWang/converter_generator/test/data"
	"github.com/jingleWang/converter_generator/test/thrifts"
)

func main() {

	cg := converter_generator.ConverterGenerator{}
	cg.OutputDir("converts")
	cg.PkgName("converts")
	cg.Alias(data.Objective{}, "ObjectiveDO")
	cg.Alias(data.LocalLifeBusinessTarget{}, "LocalLifeBusinessTargetDO")
	cg.Alias(data.ROI{}, "ROIDO")
	cg.Alias(data.OtherROI{}, "OtherROIDO")
	cg.Alias(data.LongTermROI{}, "LongTermROIDO")
	cg.Alias(data.BusinessInvestmentROI{}, "BusinessInvestmentROIDO")

	cg.Convert(data.Objective{}, thrifts.Objective{}, converter_generator.ConversionConfig{
		Name: "ConvObjective",
	})
	cg.Generate()
}
