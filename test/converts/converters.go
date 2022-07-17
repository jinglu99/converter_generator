package converts

import (
	test_data "github.com/jingleWang/converter_generator/test/data"
	test_thrifts "github.com/jingleWang/converter_generator/test/thrifts"
)

func convertMulitProfitsPtrToMulitProfitsPtr(in *test_data.MulitProfits) *test_thrifts.MulitProfits {
	if in == nil {
		return nil
	}
	out := convertMulitProfitsPtrToMulitProfits(in)
	return &out
}

func convertMusicMulitProfitPtrToMusicMulitProfit(in *test_data.MusicMulitProfit) test_thrifts.MusicMulitProfit {
	if in == nil {
		return test_thrifts.MusicMulitProfit{}
	}
	return test_thrifts.MusicMulitProfit{
		MusicCopyrightType: in.MusicCopyrightType,
		IncrementValue:     in.IncrementValue,
		GroupID:            in.GroupID,
		ClipIDs:            in.ClipIDs,
		CopyrightIncome:    in.CopyrightIncome,
	}
}

func convertProjectProfitPtrToProjectProfit(in *test_data.ProjectProfit) test_thrifts.ProjectProfit {
	if in == nil {
		return test_thrifts.ProjectProfit{}
	}
	return test_thrifts.ProjectProfit{
		ProfitType:         in.ProfitType,
		Value:              in.Value,
		ProfitTplType:      convertProfitTplTypePtrToProfitTplTypePtr(in.ProfitTplType),
		StandardizedProfit: convertStandardizedProfitPtrToStandardizedProfitPtr(in.StandardizedProfit),
		Target:             convertInt64PtrToTargetTypePtr(in.Target),
	}
}

func convertROIDOPtrToROIPtr(in *test_data.ROI) *test_thrifts.ROI {
	if in == nil {
		return nil
	}
	out := convertROIDOPtrToROI(in)
	return &out
}

func convertProfitTplTypePtrToProfitTplType(in *test_data.ProfitTplType) test_thrifts.ProfitTplType {
	if in == nil {
		return 0
	}
	return test_thrifts.ProfitTplType(*in)

}

func convertGameMulitProfitsPtrToGameMulitProfitsPtr(in *test_data.GameMulitProfits) *test_thrifts.GameMulitProfits {
	if in == nil {
		return nil
	}
	out := convertGameMulitProfitsPtrToGameMulitProfits(in)
	return &out
}

func convertChannelInfoPtrListToChannelInfoPtrList(in []*test_data.ChannelInfo) []*test_thrifts.ChannelInfo {
	if in == nil {
		return nil
	}

	out := make([]*test_thrifts.ChannelInfo, 0)
	for _, item := range in {
		out = append(out, convertChannelInfoPtrToChannelInfoPtr(item))
	}
	return out
}

func convertReturnFormulaIndexPtrToReturnFormulaIndex(in *test_data.ReturnFormulaIndex) test_thrifts.ReturnFormulaIndex {
	if in == nil {
		return test_thrifts.ReturnFormulaIndex{}
	}
	return test_thrifts.ReturnFormulaIndex{
		Key:   in.Key,
		Value: in.Value,
		Desc:  in.Desc,
	}
}

func convertReturnFormulaIndexPtrToReturnFormulaIndexPtr(in *test_data.ReturnFormulaIndex) *test_thrifts.ReturnFormulaIndex {
	if in == nil {
		return nil
	}
	out := convertReturnFormulaIndexPtrToReturnFormulaIndex(in)
	return &out
}

func convertGameMulitProfitsPtrListToGameMulitProfitsPtrList(in []*test_data.GameMulitProfits) []*test_thrifts.GameMulitProfits {
	if in == nil {
		return nil
	}

	out := make([]*test_thrifts.GameMulitProfits, 0)
	for _, item := range in {
		out = append(out, convertGameMulitProfitsPtrToGameMulitProfitsPtr(item))
	}
	return out
}

func convertLongTermROIDOPtrToLongTermROI(in *test_data.LongTermROI) test_thrifts.LongTermROI {
	if in == nil {
		return test_thrifts.LongTermROI{}
	}
	return test_thrifts.LongTermROI{
		ReturnFormulaIndexs: convertStringReturnFormulaIndexPtrMapToStringReturnFormulaIndexPtrMap(in.ReturnFormulaIndexs),
		Desc:                in.Desc,
		Roi:                 in.Roi,
		I:                   in.I,
	}
}

func convertOtherROIDOPtrToOtherROI(in *test_data.OtherROI) test_thrifts.OtherROI {
	if in == nil {
		return test_thrifts.OtherROI{}
	}
	return test_thrifts.OtherROI{
		Desc:        in.Desc,
		ROIValue:    in.ROIValue,
		CostValue:   in.CostValue,
		ReturnValue: in.ReturnValue,
	}
}

func convertGameTypePtrToGameType(in *test_data.GameType) test_thrifts.GameType {
	if in == nil {
		return 0
	}
	return test_thrifts.GameType(*in)

}

func convertProcessingMessagePtrToProcessingMessagePtr(in *test_data.ProcessingMessage) *test_thrifts.ProcessingMessage {
	if in == nil {
		return nil
	}
	out := convertProcessingMessagePtrToProcessingMessage(in)
	return &out
}

func convertStandardizedProfitPtrToStandardizedProfitPtr(in *test_data.StandardizedProfit) *test_thrifts.StandardizedProfit {
	if in == nil {
		return nil
	}
	out := convertStandardizedProfitPtrToStandardizedProfit(in)
	return &out
}

func convertContributeLevelIncomePtrToContributeLevelIncome(in *test_data.ContributeLevelIncome) test_thrifts.ContributeLevelIncome {
	if in == nil {
		return test_thrifts.ContributeLevelIncome{}
	}
	return test_thrifts.ContributeLevelIncome{
		MarkFansLevel:  convertInt64ToMarkFansLevel(in.MarkFansLevel),
		IncrementValue: in.IncrementValue,
	}
}

func convertMusicMulitProfitPtrListToMusicMulitProfitPtrList(in []*test_data.MusicMulitProfit) []*test_thrifts.MusicMulitProfit {
	if in == nil {
		return nil
	}

	out := make([]*test_thrifts.MusicMulitProfit, 0)
	for _, item := range in {
		out = append(out, convertMusicMulitProfitPtrToMusicMulitProfitPtr(item))
	}
	return out
}

func convertMovieProfitPtrToMovieProfit(in *test_data.MovieProfit) test_thrifts.MovieProfit {
	if in == nil {
		return test_thrifts.MovieProfit{}
	}
	return test_thrifts.MovieProfit{
		MovieIDs:       in.MovieIDs,
		PlatformIncome: in.PlatformIncome,
		MovieNames:     in.MovieNames,
	}
}

func convertMusicProfitInfoPtrToMusicProfitInfo(in *test_data.MusicProfitInfo) test_thrifts.MusicProfitInfo {
	if in == nil {
		return test_thrifts.MusicProfitInfo{}
	}
	return test_thrifts.MusicProfitInfo{
		MusicCopyrightType: in.MusicCopyrightType,
		Name:               in.Name,
		GroupID:            in.GroupID,
		ClipIDs:            in.ClipIDs,
	}
}

func convertMusicProfitInfoPtrListToMusicProfitInfoPtrList(in []*test_data.MusicProfitInfo) []*test_thrifts.MusicProfitInfo {
	if in == nil {
		return nil
	}

	out := make([]*test_thrifts.MusicProfitInfo, 0)
	for _, item := range in {
		out = append(out, convertMusicProfitInfoPtrToMusicProfitInfoPtr(item))
	}
	return out
}

func convertStringReturnFormulaIndexPtrMapToStringReturnFormulaIndexPtrMap(in map[string]*test_data.ReturnFormulaIndex) map[string]*test_thrifts.ReturnFormulaIndex {
	if in == nil {
		return nil
	}

	out := map[string]*test_thrifts.ReturnFormulaIndex{}
	for k, v := range in {
		key := k
		val := convertReturnFormulaIndexPtrToReturnFormulaIndexPtr(v)
		out[key] = val
	}
	return out
}

func convertOtherROIDOPtrToOtherROIPtr(in *test_data.OtherROI) *test_thrifts.OtherROI {
	if in == nil {
		return nil
	}
	out := convertOtherROIDOPtrToOtherROI(in)
	return &out
}

func convertInt64ToTargetType(in int64) test_thrifts.TargetType {

	return test_thrifts.TargetType(in)

}

func convertInt64ListToTargetTypeList(in []int64) []test_thrifts.TargetType {
	if in == nil {
		return nil
	}

	out := make([]test_thrifts.TargetType, 0)
	for _, item := range in {
		out = append(out, convertInt64ToTargetType(item))
	}
	return out
}

func convertInt64ToVerticalType(in int64) test_thrifts.VerticalType {

	return test_thrifts.VerticalType(in)

}

func convertLongTermROIDOPtrToLongTermROIPtr(in *test_data.LongTermROI) *test_thrifts.LongTermROI {
	if in == nil {
		return nil
	}
	out := convertLongTermROIDOPtrToLongTermROI(in)
	return &out
}

func convertROIDOPtrListToROIPtrList(in []*test_data.ROI) []*test_thrifts.ROI {
	if in == nil {
		return nil
	}

	out := make([]*test_thrifts.ROI, 0)
	for _, item := range in {
		out = append(out, convertROIDOPtrToROIPtr(item))
	}
	return out
}

func convertContributeLevelIncomePtrListToContributeLevelIncomePtrList(in []*test_data.ContributeLevelIncome) []*test_thrifts.ContributeLevelIncome {
	if in == nil {
		return nil
	}

	out := make([]*test_thrifts.ContributeLevelIncome, 0)
	for _, item := range in {
		out = append(out, convertContributeLevelIncomePtrToContributeLevelIncomePtr(item))
	}
	return out
}

func convertStandardizedProfitPtrToStandardizedProfit(in *test_data.StandardizedProfit) test_thrifts.StandardizedProfit {
	if in == nil {
		return test_thrifts.StandardizedProfit{}
	}
	return test_thrifts.StandardizedProfit{
		TargetType:             convertInt64PtrToTargetTypePtr(in.TargetType),
		ProfitIndex:            convertInt64PtrToProfitIndexPtr(in.ProfitIndex),
		IncrementValue:         in.IncrementValue,
		VerticalTypes:          convertInt64ListToVerticalTypeList(in.VerticalTypes),
		ExpectRate:             in.ExpectRate,
		CrowdID:                in.CrowdID,
		CrowdEstimatedNum:      in.CrowdEstimatedNum,
		CrowdConfigJSON:        in.CrowdConfigJSON,
		CrowdType:              in.CrowdType,
		TargetCustomName:       in.TargetCustomName,
		ProfitIndexName:        in.ProfitIndexName,
		CrowdUpdateTime:        in.CrowdUpdateTime,
		GameIDs:                in.GameIDs,
		CustomValue:            in.CustomValue,
		TopicIDs:               in.TopicIDs,
		MulitProfits:           convertMulitProfitsPtrListToMulitProfitsPtrList(in.MulitProfits),
		AdditionalRemarks:      in.AdditionalRemarks,
		ContributeLevelIncomes: convertContributeLevelIncomePtrListToContributeLevelIncomePtrList(in.ContributeLevelIncomes),
		ActivityHotSpots:       in.ActivityHotSpots,
		GameMulitProfits:       convertGameMulitProfitsPtrListToGameMulitProfitsPtrList(in.GameMulitProfits),
		ProcessingMessage:      convertProcessingMessagePtrToProcessingMessagePtr(in.ProcessingMessage),
		MusicMulitProfits:      convertMusicMulitProfitPtrListToMusicMulitProfitPtrList(in.MusicMulitProfits),
		MovieProfit:            convertMovieProfitPtrToMovieProfitPtr(in.MovieProfit),
		ActivityHotSpotIDs:     in.ActivityHotSpotIDs,
		MusicProfitInfos:       convertMusicProfitInfoPtrListToMusicProfitInfoPtrList(in.MusicProfitInfos),
		AbilityName:            in.AbilityName,
		LibraURL:               in.LibraURL,
		ChannelInfoList:        convertChannelInfoPtrListToChannelInfoPtrList(in.ChannelInfoList),
	}
}

func convertROIDOPtrToROI(in *test_data.ROI) test_thrifts.ROI {
	if in == nil {
		return test_thrifts.ROI{}
	}
	return test_thrifts.ROI{
		Value:                 in.Value,
		Desc:                  in.Desc,
		ROIType:               convertROITypePtrToROITypePtr(in.ROIType),
		ShortTermROI:          convertShortTermROIPtrToShortTermROIPtr(in.ShortTermROI),
		BusinessInvestmentROI: convertBusinessInvestmentROIDOPtrToBusinessInvestmentROIPtr(in.BusinessInvestmentROI),
		LongTermROI:           convertLongTermROIDOPtrToLongTermROIPtr(in.LongTermROI),
		OtherROI:              convertOtherROIDOPtrToOtherROIPtr(in.OtherROI),
	}
}

func convertLocalLifeBusinessTargetDOPtrToLocalLifeBusinessTarget(in *test_data.LocalLifeBusinessTarget) test_thrifts.LocalLifeBusinessTarget {
	if in == nil {
		return test_thrifts.LocalLifeBusinessTarget{}
	}
	return test_thrifts.LocalLifeBusinessTarget{
		TargetID:    in.TargetID,
		TargetVlue:  in.TargetVlue,
		CustomValue: in.CustomValue,
	}
}

func ConvObjective(in test_data.Objective) test_thrifts.Objective {
	return test_thrifts.Objective{
		Population:               in.Population,
		Targets:                  convertInt64ListToTargetTypeList(in.Targets),
		Profits:                  convertProjectProfitPtrListToProjectProfitPtrList(in.Profits),
		ROI:                      convertROIDOPtrToROIPtr(in.ROI),
		ParentBudgetDepartmentID: in.ParentBudgetDepartmentID,
		ROIs:                     convertROIDOPtrListToROIPtrList(in.ROIs),
		MainROIType:              convertROITypePtrToROITypePtr(in.MainROIType),
		LocalLifeBusinessTargets: convertLocalLifeBusinessTargetDOPtrListToLocalLifeBusinessTargetPtrList(in.LocalLifeBusinessTargets),
	}
}

func convertMusicProfitInfoPtrToMusicProfitInfoPtr(in *test_data.MusicProfitInfo) *test_thrifts.MusicProfitInfo {
	if in == nil {
		return nil
	}
	out := convertMusicProfitInfoPtrToMusicProfitInfo(in)
	return &out
}

func convertInt64PtrToProfitIndexPtr(in *int64) *test_thrifts.ProfitIndex {
	if in == nil {
		return nil
	}
	out := convertInt64PtrToProfitIndex(in)
	return &out
}

func convertMulitProfitsPtrListToMulitProfitsPtrList(in []*test_data.MulitProfits) []*test_thrifts.MulitProfits {
	if in == nil {
		return nil
	}

	out := make([]*test_thrifts.MulitProfits, 0)
	for _, item := range in {
		out = append(out, convertMulitProfitsPtrToMulitProfitsPtr(item))
	}
	return out
}

func convertProjectProfitPtrToProjectProfitPtr(in *test_data.ProjectProfit) *test_thrifts.ProjectProfit {
	if in == nil {
		return nil
	}
	out := convertProjectProfitPtrToProjectProfit(in)
	return &out
}

func convertBusinessInvestmentROIDOPtrToBusinessInvestmentROI(in *test_data.BusinessInvestmentROI) test_thrifts.BusinessInvestmentROI {
	if in == nil {
		return test_thrifts.BusinessInvestmentROI{}
	}
	return test_thrifts.BusinessInvestmentROI{
		ReturnFormulaIndexs: convertStringReturnFormulaIndexPtrMapToStringReturnFormulaIndexPtrMap(in.ReturnFormulaIndexs),
		Desc:                in.Desc,
		Roi:                 in.Roi,
		I:                   in.I,
	}
}

func convertBusinessInvestmentROIDOPtrToBusinessInvestmentROIPtr(in *test_data.BusinessInvestmentROI) *test_thrifts.BusinessInvestmentROI {
	if in == nil {
		return nil
	}
	out := convertBusinessInvestmentROIDOPtrToBusinessInvestmentROI(in)
	return &out
}

func convertProfitTplTypePtrToProfitTplTypePtr(in *test_data.ProfitTplType) *test_thrifts.ProfitTplType {
	if in == nil {
		return nil
	}
	out := convertProfitTplTypePtrToProfitTplType(in)
	return &out
}

func convertContributeLevelIncomePtrToContributeLevelIncomePtr(in *test_data.ContributeLevelIncome) *test_thrifts.ContributeLevelIncome {
	if in == nil {
		return nil
	}
	out := convertContributeLevelIncomePtrToContributeLevelIncome(in)
	return &out
}

func convertROITypePtrToROITypePtr(in *test_data.ROIType) *test_thrifts.ROIType {
	if in == nil {
		return nil
	}
	out := convertROITypePtrToROIType(in)
	return &out
}

func convertLocalLifeBusinessTargetDOPtrToLocalLifeBusinessTargetPtr(in *test_data.LocalLifeBusinessTarget) *test_thrifts.LocalLifeBusinessTarget {
	if in == nil {
		return nil
	}
	out := convertLocalLifeBusinessTargetDOPtrToLocalLifeBusinessTarget(in)
	return &out
}

func convertInt64PtrToProfitIndex(in *int64) test_thrifts.ProfitIndex {
	if in == nil {
		return 0
	}
	return test_thrifts.ProfitIndex(*in)

}

func convertInt64ListToVerticalTypeList(in []int64) []test_thrifts.VerticalType {
	if in == nil {
		return nil
	}

	out := make([]test_thrifts.VerticalType, 0)
	for _, item := range in {
		out = append(out, convertInt64ToVerticalType(item))
	}
	return out
}

func convertMulitProfitsPtrToMulitProfits(in *test_data.MulitProfits) test_thrifts.MulitProfits {
	if in == nil {
		return test_thrifts.MulitProfits{}
	}
	return test_thrifts.MulitProfits{
		ProfitIndex:      convertInt64PtrToProfitIndexPtr(in.ProfitIndex),
		IncrementValue:   in.IncrementValue,
		CustomValue:      in.CustomValue,
		ExpectRate:       in.ExpectRate,
		ProfitIndexName:  in.ProfitIndexName,
		TargetCustomName: in.TargetCustomName,
	}
}

func convertMusicMulitProfitPtrToMusicMulitProfitPtr(in *test_data.MusicMulitProfit) *test_thrifts.MusicMulitProfit {
	if in == nil {
		return nil
	}
	out := convertMusicMulitProfitPtrToMusicMulitProfit(in)
	return &out
}

func convertLocalLifeBusinessTargetDOPtrListToLocalLifeBusinessTargetPtrList(in []*test_data.LocalLifeBusinessTarget) []*test_thrifts.LocalLifeBusinessTarget {
	if in == nil {
		return nil
	}

	out := make([]*test_thrifts.LocalLifeBusinessTarget, 0)
	for _, item := range in {
		out = append(out, convertLocalLifeBusinessTargetDOPtrToLocalLifeBusinessTargetPtr(item))
	}
	return out
}

func convertMovieProfitPtrToMovieProfitPtr(in *test_data.MovieProfit) *test_thrifts.MovieProfit {
	if in == nil {
		return nil
	}
	out := convertMovieProfitPtrToMovieProfit(in)
	return &out
}

func convertChannelInfoPtrToChannelInfo(in *test_data.ChannelInfo) test_thrifts.ChannelInfo {
	if in == nil {
		return test_thrifts.ChannelInfo{}
	}
	return test_thrifts.ChannelInfo{
		ChannelID:   in.ChannelID,
		ChannelName: in.ChannelName,
	}
}

func convertROITypePtrToROIType(in *test_data.ROIType) test_thrifts.ROIType {
	if in == nil {
		return 0
	}
	return test_thrifts.ROIType(*in)

}

func convertShortTermROIPtrToShortTermROIPtr(in *test_data.ShortTermROI) *test_thrifts.ShortTermROI {
	if in == nil {
		return nil
	}
	out := convertShortTermROIPtrToShortTermROI(in)
	return &out
}

func convertInt64PtrToTargetTypePtr(in *int64) *test_thrifts.TargetType {
	if in == nil {
		return nil
	}
	out := convertInt64PtrToTargetType(in)
	return &out
}

func convertInt64ToMarkFansLevel(in int64) test_thrifts.MarkFansLevel {

	return test_thrifts.MarkFansLevel(in)

}

func convertGameTypePtrToGameTypePtr(in *test_data.GameType) *test_thrifts.GameType {
	if in == nil {
		return nil
	}
	out := convertGameTypePtrToGameType(in)
	return &out
}

func convertGameMulitProfitsPtrToGameMulitProfits(in *test_data.GameMulitProfits) test_thrifts.GameMulitProfits {
	if in == nil {
		return test_thrifts.GameMulitProfits{}
	}
	return test_thrifts.GameMulitProfits{
		GameID:       in.GameID,
		GameName:     in.GameName,
		MulitProfits: convertMulitProfitsPtrListToMulitProfitsPtrList(in.MulitProfits),
		LTV:          in.LTV,
		GameType:     convertGameTypePtrToGameTypePtr(in.GameType),
	}
}

func convertChannelInfoPtrToChannelInfoPtr(in *test_data.ChannelInfo) *test_thrifts.ChannelInfo {
	if in == nil {
		return nil
	}
	out := convertChannelInfoPtrToChannelInfo(in)
	return &out
}

func convertShortTermROIPtrToShortTermROI(in *test_data.ShortTermROI) test_thrifts.ShortTermROI {
	if in == nil {
		return test_thrifts.ShortTermROI{}
	}
	return test_thrifts.ShortTermROI{
		ReturnFormulaIndexs: convertStringReturnFormulaIndexPtrMapToStringReturnFormulaIndexPtrMap(in.ReturnFormulaIndexs),
		Desc:                in.Desc,
		Roi:                 in.Roi,
		I:                   in.I,
	}
}

func convertInt64PtrToTargetType(in *int64) test_thrifts.TargetType {
	if in == nil {
		return 0
	}
	return test_thrifts.TargetType(*in)

}

func convertProcessingMessagePtrToProcessingMessage(in *test_data.ProcessingMessage) test_thrifts.ProcessingMessage {
	if in == nil {
		return test_thrifts.ProcessingMessage{}
	}
	return test_thrifts.ProcessingMessage{
		GameID:          in.GameID,
		GameName:        in.GameName,
		LTV:             in.LTV,
		GameType:        convertGameTypePtrToGameTypePtr(in.GameType),
		GameRolePercent: in.GameRolePercent,
	}
}

func convertProjectProfitPtrListToProjectProfitPtrList(in []*test_data.ProjectProfit) []*test_thrifts.ProjectProfit {
	if in == nil {
		return nil
	}

	out := make([]*test_thrifts.ProjectProfit, 0)
	for _, item := range in {
		out = append(out, convertProjectProfitPtrToProjectProfitPtr(item))
	}
	return out
}
